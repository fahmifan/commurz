package webserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fahmifan/authme/backend/httphandler"
	"github.com/fahmifan/commurz/pkg/auth"
	"github.com/fahmifan/commurz/pkg/config"
	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/fahmifan/commurz/pkg/pb/commurz/v1/commurzv1connect"
	"github.com/fahmifan/commurz/pkg/service"
	"github.com/fahmifan/commurz/pkg/web/webserver/httperr"
	"github.com/fahmifan/flycasbin/acl"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	vueglue "github.com/torenware/vite-go"
)

type Webserver struct {
	cookieAutherHandler *httphandler.CookieAuthHandler
	echo                *echo.Echo
	service             *service.Service
	port                int
}

func NewWebserver(service *service.Service, port int, cookieAutherHandler *httphandler.CookieAuthHandler) *Webserver {
	return &Webserver{
		service:             service,
		port:                port,
		cookieAutherHandler: cookieAutherHandler,
	}
}

func (server *Webserver) Stop(ctx context.Context) error {
	return server.echo.Shutdown(ctx)
}

func (server *Webserver) Run() error {
	server.echo = echo.New()

	pageMdw := PageMiddleware{server}
	// apiMdw := APIMiddleware{server}

	server.echo.Renderer = &Renderer{}
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

	authRouter, err := server.cookieAutherHandler.CookieAuthRouter()
	if err != nil {
		return fmt.Errorf("webserver: cookieAutherHandler.CookieAuthRouter: %w", err)
	}

	server.cookieAutherHandler.SetRedirectAfterLogin(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("redirect after login")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	cookieMdw := server.cookieAutherHandler.Middleware()

	server.echo.Use(
		echo.WrapMiddleware(cookieMdw.SetAuthUserToCtx()),
	)

	if err = server.routeFE(server.echo, &pageMdw); err != nil {
		return fmt.Errorf("webserver: routeFE: %w", err)
	}

	server.echo.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	server.echo.Any("/api/auth*", echo.WrapHandler(authRouter), trimPath("*"))
	server.echo.Group("/grpc").Any(
		grpHandlerName+"*",
		echo.WrapHandler(grpcHandler),
		echo.WrapMiddleware(cookieMdw.CSRF()),
		trimPath("/grpc"),
	)

	return server.echo.Start(server.getPort())
}

func (server *Webserver) getPort() string {
	if server.port <= 0 {
		return ":8080"
	}

	return ":" + fmt.Sprint(server.port)
}

func (s *Webserver) routeAllFEs(group *echo.Group, hh echo.HandlerFunc, pageMdw *PageMiddleware) {
	group.GET("/", hh).Name = "page-index" // login

	group.GET("/backoffice/products", hh, pageMdw.HasAccess([]service.Perm{
		Perm(auth.Manage, auth.Product),
	})).Name = "page-backoffice-products"

	group.GET("/backoffice/products/stocks", hh, pageMdw.HasAccess([]service.Perm{
		Perm(auth.Manage, auth.Product),
	})).Name = "page-backoffice-products-stocks"
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

	viteRenderer := s.renderVite("templates/vite/index.html", viteGlue)
	reactGroup := s.echo.Group("")
	{
		s.routeAllFEs(reactGroup, viteRenderer, pageMdw)
	}
	return nil
}

func (s *Webserver) renderVite(tplName string, viteGlue *vueglue.VueGlue) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, tplName, echo.Map{"vite": viteGlue})
	}
}

func Perm(action acl.Action, resource acl.Resource) service.Perm {
	return service.Perm{Action: action, Resource: resource}
}

type PageMiddleware struct {
	*Webserver
}

func (mw *PageMiddleware) HasAccess(perms []service.Perm) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			authUser, ok := httphandler.GetUser(ec.Request().Context())
			if !ok {
				return ec.Redirect(http.StatusFound, "/api/auth/new")
			}

			userID, err := uuid.Parse(authUser.GUID)
			if err != nil {
				return ec.Redirect(http.StatusFound, "/api/auth/new")
			}

			err = mw.service.InternalHasAccess(ec.Request().Context(), userID, perms)
			if err != nil {
				if !errors.Is(err, acl.ErrPermissionDenied) {
					logs.ErrCtx(ec.Request().Context(), err, "[HasAccess] mw.acl.Can")
					return ec.JSON(http.StatusInternalServerError, echo.Map{
						"error": "internal error",
					})
				}

				return ec.JSON(http.StatusForbidden, echo.Map{
					"error": err.Error(),
				})
			}

			return next(ec)
		}
	}
}

func (mw *PageMiddleware) MustAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			_, ok := httphandler.GetUser(ec.Request().Context())
			if !ok {
				return ec.Redirect(http.StatusFound, ec.Echo().Reverse("page-auth-login-new"))
			}

			return next(ec)
		}
	}
}

func (mw *PageMiddleware) MustNonAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			_, ok := httphandler.GetUser(ec.Request().Context())
			if ok {
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
			_, ok := httphandler.GetUser(ec.Request().Context())
			if !ok {
				return httperr.JSON(ec, httperr.ErrUnauthorized)
			}

			return next(ec)
		}
	}
}

func trimPath(groupPrefix string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().URL.Path = strings.TrimPrefix(c.Request().URL.Path, groupPrefix)
			return next(c)
		}
	}
}
