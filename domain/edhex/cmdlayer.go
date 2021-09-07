package edhex

import (
	"fmt"
	"path/filepath"

	"github.com/edteamlat/go-wizard/model"
)

const CMDLayerName = "cmd"

const cmdFolder = "cmd"

var cmdInitActionTemplates = model.Templates{
	{
		Name:     "config.gotpl",
		Filename: "config.go",
		Path:     cmdFolder,
	},
	{
		Name:     "database.gotpl",
		Filename: "database.go",
		Path:     cmdFolder,
	},
	{
		Name:     "echo.gotpl",
		Filename: "echo.go",
		Path:     cmdFolder,
	},
	{
		Name:     "logger.gotpl",
		Filename: "logger.go",
		Path:     cmdFolder,
	},
	{
		Name:     "main.gotpl",
		Filename: "main.go",
		Path:     cmdFolder,
	},
	{
		Name:     "remoteconfig.gotpl",
		Filename: "remoteconfig.go",
		Path:     cmdFolder,
	},
	{
		Name:     "configuration-example.gotpl",
		Filename: "configuration.json",
		Path:     cmdFolder,
	},
}

type cmdLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewCMDLayer(template UseCaseTemplate, storage Storage) cmdLayer {
	return cmdLayer{template: template, storage: storage}
}

func (d cmdLayer) Init(data model.Layer) error {
	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "cmd")); err != nil {
		return fmt.Errorf("edhex-cmdlayer: %w", err)
	}

	if err := bulkTemplates(d.template, d.storage, cmdInitActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-cmdlayer: %w", err)
	}

	return nil
}

func (d cmdLayer) Create(data model.Layer) error { return nil }

func (d cmdLayer) Override(m model.Layer) error { return nil }

func (d cmdLayer) AddField(m model.Layer) error { return nil }
