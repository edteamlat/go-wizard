package edhex

import (
	"fmt"
	"os"
	"os/exec"
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
	{
		Name:     "wizard.gotpl",
		Filename: "wizard-config.yaml",
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

	if err := d.runCommands(data); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	return nil
}

func (d rootLayer) runCommands(data model.Layer) error {
	if err := d.CDToProject(data.GetProjectName()); err != nil {
		return err
	}

	if err := d.initGit(); err != nil {
		return err
	}

	if err := d.initGoMod(data.ModuleName); err != nil {
		return err
	}

	return nil
}

func (d rootLayer) CDToProject(projectName string) error {
	return os.Chdir(fmt.Sprintf("./%s", projectName))
}

func (d rootLayer) initGoMod(moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)

	if err := cmd.Run(); err != nil {
		return err
	}

	tidyCMD := exec.Command("go", "mod", "tidy")

	return tidyCMD.Run()
}

func (d rootLayer) initGit() error {
	cmd := exec.Command("git", "init")

	return cmd.Run()
}

func (d rootLayer) Create(data model.Layer) error { return nil }

func (d rootLayer) Override(m model.Layer) error { return nil }

func (d rootLayer) AddField(m model.Layer) error { return nil }
