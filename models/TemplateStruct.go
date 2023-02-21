package models

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t.templates == nil {
		return fmt.Errorf("template is nil")
	}
	if name == "" {
		return fmt.Errorf("template name is empty")
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	if data == nil {
		return fmt.Errorf("data is nil")
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate(tpl *template.Template) *Template {
	return &Template{templates: tpl}
}
