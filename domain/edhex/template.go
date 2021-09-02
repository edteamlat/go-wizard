package edhex

import (
	"bytes"

	"github.com/edteamlat/go-wizard/model"
)

func createTemplate(template UseCaseTemplate, storage Storage, data model.Template) error {
	fileBuf := bytes.Buffer{}
	if err := template.Create(&fileBuf, cmdConfigTemplateName, data.Layer); err != nil {
		return err
	}

	if err := storage.Save(data.Path, fileBuf); err != nil {
		return err
	}

	return nil
}
