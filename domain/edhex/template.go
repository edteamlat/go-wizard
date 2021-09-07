package edhex

import (
	"bytes"

	"github.com/edteamlat/go-wizard/model"
)

func bulkTemplates(template UseCaseTemplate, storage Storage, templates model.Templates, data model.Layer) error {
	for _, v := range templates {
		v.SetPathPrefix(data)
		v.SetLayerData(data)

		if err := createTemplate(template, storage, v); err != nil {
			return err
		}
	}

	return nil
}

func createTemplate(template UseCaseTemplate, storage Storage, data model.Template) error {
	fileBuf := bytes.Buffer{}
	if err := template.Create(&fileBuf, data.Name, data.Layer); err != nil {
		return err
	}

	if err := storage.Save(data.Path, fileBuf); err != nil {
		return err
	}

	return nil
}
