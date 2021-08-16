package texttemplate

import (
	"fmt"
	"io"
	"text/template"

	"github.com/edteamlat/go-wizard/model"
)

type textTemplate struct {
	tpl *template.Template
}

func NewTemplate(tpl *template.Template) *textTemplate {
	return &textTemplate{tpl: tpl}
}

func (t textTemplate) Create(wr io.Writer, templateName string, data model.Layer) error {
	if err := t.tpl.ExecuteTemplate(wr, templateName, data); err != nil {
		return fmt.Errorf("template: %w", err)
	}

	return nil
}
