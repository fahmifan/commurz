package webserver

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fahmifan/commurz/pkg/config"
	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/fahmifan/commurz/pkg/pb/commurz/v1/commurzv1connect"
	"github.com/fahmifan/commurz/pkg/service"
	"github.com/fahmifan/commurz/pkg/web/webserver/httperr"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	vueglue "github.com/torenware/vite-go"
)

type Webserver struct {
	db      *sql.DB
	echo    *echo.Echo
	service *service.Service
	port    int
	session *Session
}

func NewWebserver(service *service.Service, port int) *Webserver {
	return &Webserver{
		service: service,
		port:    port,
	}
}

func (server *Webserver) Run() error {
	server.echo = echo.New()

	pageMdw := PageMiddleware{server}
	apiMdw := APIMiddleware{server}

	server.echo.Use(
		middleware.RemoveTrailingSlash(),
		logs.EchoRequestID(),
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogValuesFunc: logs.EchoRequestLogger(true),
			LogLatency:    true,
			LogRemoteIP:   true,
			LogUserAgent:  true,
			LogError:      true,
			HandleError:   true,
		}),
	)

	grpHandlerName, grpcHandler := commurzv1connect.NewCommurzServiceHandler(
		server.service,
	)
	server.echo.Group("/grpc").Any(
		grpHandlerName+"*",
		echo.WrapHandler(grpcHandler),
		apiMdw.MustAuth(),
		trimPathGroup("/grpc"),
	)

	server.routeFE(server.echo, &pageMdw)

	return server.echo.Start(server.getPort())
}

func trimPathGroup(groupPrefix string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().URL.Path = strings.TrimPrefix(c.Request().URL.Path, groupPrefix)
			return next(c)
		}
	}
}

func (server *Webserver) getPort() string {
	if server.port <= 0 {
		return ":8080"
	}

	return ":" + fmt.Sprint(server.port)
}

func (server *Webserver) Stop(ctx context.Context) error {
	return server.echo.Shutdown(ctx)
}

func (s *Webserver) routeFE(ec *echo.Echo, pageMdw *PageMiddleware) error {
	return s.routeDevFE(ec, pageMdw)
}

func (s *Webserver) routeDevFE(ec *echo.Echo, pageMdw *PageMiddleware) (err error) {
	viteCfg := &vueglue.ViteConfig{
		Environment:   config.Env(),
		JSProjectPath: "frontend",
		EntryPoint:    "src/main.tsx",
		Platform:      "react",
		FS:            os.DirFS("../frontend"),
		DevServerPort: "3000",
	}

	err = viteCfg.SetDevelopmentDefaults()
	if err != nil {
		fmt.Println("viteCfg.SetDevelopmentDefaults() error:", err)
		return err
	}

	viteGlue, err := vueglue.NewVueGlue(viteCfg)
	if err != nil {
		return err
	}
	viteGlue.Debug = true

	viteFileSrv, err := viteGlue.FileServer()
	if err != nil {
		return err
	}

	staticGroup := s.echo.Group("")
	{
		staticGroup.GET("/assets/*", echo.WrapHandler(viteFileSrv))
		// staticGroup.StaticFS("/favicons", faviconFS)
	}

	viteRenderer := s.renderVite("pages/vite/index.html", viteGlue)
	reactGroup := s.echo.Group("")
	{
		s.renderAllFEs(reactGroup, viteRenderer, pageMdw)
	}
	return nil
}

func (s *Webserver) renderAllFEs(group *echo.Group, hh echo.HandlerFunc, pageMdw *PageMiddleware) {
	group.GET("/", hh, pageMdw.MustAuth()).Name = "page-index"
	group.GET("/journals", hh, pageMdw.MustAuth()).Name = "page-carts"
	group.GET("/carts/*", hh, pageMdw.MustAuth()).Name = "page-carts"
	group.GET("/auth/login/new", hh, pageMdw.MustNonAuth()).Name = "page-auth-login-new"
	group.GET("/auth/register/new", hh, pageMdw.MustNonAuth()).Name = "page-auth-register-new"
}

func (s *Webserver) renderVite(tplName string, viteGlue *vueglue.VueGlue) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, tplName, echo.Map{"vite": viteGlue})
	}
}

type PageMiddleware struct {
	*Webserver
}

func (mw *PageMiddleware) MustAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			authUser := mw.session.GetUser(ec)
			if authUser == nil {
				return ec.Redirect(http.StatusFound, ec.Echo().Reverse("page-auth-login-new"))
			}

			return next(ec)
		}
	}
}

func (mw *PageMiddleware) MustNonAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			authUser := mw.session.GetUser(ec)
			if authUser != nil {
				return ec.Redirect(http.StatusFound, "/")
			}

			return next(ec)
		}
	}
}

type APIMiddleware struct {
	*Webserver
}

func (mw *APIMiddleware) MustAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			authUser := mw.session.GetUser(ec)
			if authUser == nil {
				return httperr.JSON(ec, httperr.ErrUnauthorized)
			}

			ctx := service.CtxWithUser(ec.Request().Context(), service.User{
				ID: authUser.UserID,
			})

			ec.SetRequest(ec.Request().WithContext(ctx))

			return next(ec)
		}
	}
}
