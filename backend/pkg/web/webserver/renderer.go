package webserver

import (
	"io"

	"github.com/fahmifan/commurz/pkg/web"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

type Renderer struct{}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	feLoader := pongo2.NewSet("feLoader", pongo2.NewFSLoader(web.TemplateFS()))

	tpl, err := feLoader.FromFile("templates/vite/index.html")
	if err != nil {
		return err
	}

	var ctx pongo2.Context
	if data != nil {
		switch val := data.(type) {
		case pongo2.Context:
			ctx = pongo2.Context(val)
		case echo.Map:
			ctx = pongo2.Context(val)
		case map[string]interface{}:
			ctx = pongo2.Context(val)
		default:
			ctx = make(pongo2.Context)
		}
	} else {
		ctx = make(pongo2.Context)
	}

	return tpl.ExecuteWriter(ctx, w)
}
