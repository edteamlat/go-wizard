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
		Filename: "configuration.example.json",
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
	if err := d.createDirs(data); err != nil {
		if err != nil {
			return fmt.Errorf("edhex-cmdlayer: %w", err)
		}
	}

	if err := bulkFromTemplates(d.template, d.storage, cmdInitActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-cmdlayer: %w", err)
	}

	return nil
}

func (d cmdLayer) createDirs(data model.Layer) error {
	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "cmd")); err != nil {
		return err
	}

	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "cmd", "certificates")); err != nil {
		return err
	}

	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "cmd", "logs")); err != nil {
		return err
	}

	return nil
}

func (d cmdLayer) SuccessfulMsg(prefixCount string) {
	fmt.Printf("%s cmd layer generated ✅\n", prefixCount)
}

func (d cmdLayer) FailureMsg(prefixCount string, err error) {
	fmt.Printf("%s cmd layer failed 🚨, %s\n", prefixCount, err.Error())
}
