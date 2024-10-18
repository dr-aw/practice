package httpHandler

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"path/filepath"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initializeTemplates() (*TemplateRenderer, error) {
	templatesPath := filepath.Join("templates", "*.html")
	templates, err := template.ParseGlob(templatesPath)
	if err != nil {
		return nil, err
	}
	return &TemplateRenderer{
		templates: templates,
	}, nil
}
