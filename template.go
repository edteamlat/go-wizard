package main

import (
	"fmt"
	"io"
	"text/template"
)

type UseCaseTemplate interface {
	Create(wr io.Writer, templateName string, data interface{}) error
}

type Template struct {
	tpl *template.Template
}

func NewTemplate(tpl *template.Template) *Template {
	return &Template{tpl: tpl}
}

func (t Template) Create(wr io.Writer, templateName string, data interface{}) error {
	if err := t.tpl.ExecuteTemplate(wr, templateName, data); err != nil {
		return fmt.Errorf("template: %w", err)
	}

	return nil
}
