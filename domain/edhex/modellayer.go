package edhex

import (
	"fmt"
	"path/filepath"

	"github.com/edteamlat/go-wizard/model"
)

const (
	modelTemplateName = "newmodel.gotpl"
)

const ModelLayerName = "model"
const modelFolder = "model"

var (
	modelInitActionTemplates = model.Templates{
		{
			Name:     "error.gotpl",
			Filename: "error.go",
		},
		{
			Name:     "model_test.gotpl",
			Filename: "model_test.go",
		},
		{
			Name:     "model.gotpl",
			Filename: "model.go",
		},
		{
			Name:     "modelconfig.gotpl",
			Filename: "config.gotpl",
		},
		{
			Name:     "modellogger.gotpl",
			Filename: "logger.go",
		},
		{
			Name:     "modelremoteconfig.gotpl",
			Filename: "remoteconfig.go",
		},
		{
			Name:     "modelrouter.gotpl",
			Filename: "router.go",
		},
	}

	modelAddActionTemplates = model.Templates{
		{
			Name:     "newmodel.gotpl",
			Filename: "%s.go", // the name will be the name of the package
			Path:     modelFolder,
		},
	}
)

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

	modelInitActionTemplates.SetPath(modelFolder)
	if err := bulkTemplates(d.template, d.storage, modelInitActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	return nil
}

func (d modelLayer) Create(data model.Layer) error {
	if err := bulkTemplates(d.template, d.storage, modelAddActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	return nil
}

func (d modelLayer) Override(m model.Layer) error {
	return nil
}

func (d modelLayer) AddField(m model.Layer) error {
	return nil
}
