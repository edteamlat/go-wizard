package main

import (
	"fmt"
	"io"
	"text/template"
)

type useCaseTemplate interface {
	create(wr io.Writer, templateName string, data layer) error
}

type textTemplate struct {
	tpl *template.Template
}

func newTemplate(tpl *template.Template) *textTemplate {
	return &textTemplate{tpl: tpl}
}

func (t textTemplate) create(wr io.Writer, templateName string, data layer) error {
	if err := t.tpl.ExecuteTemplate(wr, templateName, data); err != nil {
		return fmt.Errorf("template: %w", err)
	}

	return nil
}
