package edhex

import (
	"fmt"
	"path/filepath"

	"github.com/edteamlat/go-wizard/model"
)

const RootLayerName = "root"

var rootInitActionTemplates = model.Templates{
	{
		Name:     "editorconfig.gotpl",
		Filename: ".editorconfig",
	},
	{
		Name:     "gitignore.gotpl",
		Filename: ".gitignore",
	},
	{
		Name:     "readme.gotpl",
		Filename: "README.md",
	},
}

type rootLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewRootLayer(template UseCaseTemplate, storage Storage) rootLayer {
	return rootLayer{template: template, storage: storage}
}

func (d rootLayer) Init(data model.Layer) error {
	if err := d.storage.CreateDir(data.ProjectPath); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "infrastructure")); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	if err := bulkTemplates(d.template, d.storage, rootInitActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	return nil
}

func (d rootLayer) Create(data model.Layer) error { return nil }

func (d rootLayer) Override(m model.Layer) error { return nil }

func (d rootLayer) AddField(m model.Layer) error { return nil }
