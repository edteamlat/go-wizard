package edhex

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/edteamlat/go-wizard/model"
)

const (
	modelTemplateName = "newmodel.gotpl"
)

const ModelLayerName = "model"

type modelLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewModelLayer(template UseCaseTemplate, storage Storage) modelLayer {
	return modelLayer{template: template, storage: storage}
}

func (d modelLayer) Init(data model.Layer) error {
	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "model")); err != nil {
		return fmt.Errorf("edhex-modellayer: %w", err)
	}

	return nil
}

func (d modelLayer) Create(data model.Layer) error {
	if err := d.createNewModel(data); err != nil {
		return fmt.Errorf("edhex-modellayer: %w", err)
	}

	return nil
}

func (d modelLayer) createNewModel(data model.Layer) error {
	filename := fmt.Sprintf("%s.go", strings.ToLower(data.Model))

	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  modelTemplateName,
		Path:  data.GetPath(ModelLayerName, filename, false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d modelLayer) Override(m model.Layer) error {
	return nil
}

func (d modelLayer) AddField(m model.Layer) error {
	return nil
}
