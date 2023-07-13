package webserver

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/fahmifan/commurz/pkg/service"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

type Webserver struct {
	echo    *echo.Echo
	service *service.Service
	port    int
}

func NewWebserver(service *service.Service, port int) *Webserver {
	return &Webserver{
		service: service,
		port:    port,
	}
}

func (server *Webserver) Run() error {
	server.echo = echo.New()
	server.echo.Renderer = &renderer{rootDir: "pkg/web/templates"}

	homeHandler := HomeHandler{server}
	userHandler := UserHandler{server}

	server.echo.GET("/", homeHandler.Index())
	server.echo.POST("/users", userHandler.Create())

	return server.echo.Start(server.getPort())
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

type renderer struct {
	rootDir string
}

func (r *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tpl, err := pongo2.FromFile(path.Join(r.rootDir, name))
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	return tpl.ExecuteWriter(pongo2.Context{"data": data}, w)
}
