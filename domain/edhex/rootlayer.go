package edhex

import (
	"fmt"
	"path/filepath"

	"github.com/edteamlat/go-wizard/model"
)

const (
	editorConfigTemplateName = "editorconfig.gotpl"
	gitignoreTemplateName    = "gitignore.gotpl"
	readmeTemplateName       = "readme.gotpl"
)

const RootLayerName = "root"

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

	if err := d.createEditorConfig(data); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	if err := d.createGitignore(data); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	if err := d.createREADME(data); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	return nil
}

func (d rootLayer) createEditorConfig(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  editorConfigTemplateName,
		Path:  data.GetPath("", ".editorconfig", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d rootLayer) createGitignore(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  gitignoreTemplateName,
		Path:  data.GetPath("", ".gitignore", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d rootLayer) createREADME(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  readmeTemplateName,
		Path:  data.GetPath("", "README.md", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d rootLayer) Create(data model.Layer) error { return nil }

func (d rootLayer) Override(m model.Layer) error { return nil }

func (d rootLayer) AddField(m model.Layer) error { return nil }
