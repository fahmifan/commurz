package web

import "embed"

//go:embed templates
var templateFS embed.FS

func TemplateFS() embed.FS {
	return templateFS
}
