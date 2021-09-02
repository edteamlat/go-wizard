package edhex

import (
	"bytes"
	"fmt"

	"github.com/edteamlat/go-wizard/model"
)

const (
	editorConfigTemplateName = "editorconfig.gotpl"
	gitignoreTemplateName    = "gitignore.gotpl"
)

const RootLayerName = "root"

type rootLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewRootLayer(template UseCaseTemplate, storage Storage) rootLayer {
	return rootLayer{template: template, storage: storage}
}

func (d rootLayer) Init(m model.Layer) error {
	if err := d.createEditorConfig(m); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	if err := d.createGitignore(m); err != nil {
		return fmt.Errorf("edhex-rootlayer: %w", err)
	}

	return nil
}

func (d rootLayer) createEditorConfig(data model.Layer) error {
	fileBuf := bytes.Buffer{}
	if err := d.template.Create(&fileBuf, editorConfigTemplateName, data); err != nil {
		return err
	}

	if err := d.storage.Save(data.GetPath("", ".editorconfig", false), fileBuf); err != nil {
		return err
	}

	return nil
}

func (d rootLayer) createGitignore(data model.Layer) error {
	fileBuf := bytes.Buffer{}
	if err := d.template.Create(&fileBuf, gitignoreTemplateName, data); err != nil {
		return err
	}

	if err := d.storage.Save(data.GetPath("", ".gitignore", false), fileBuf); err != nil {
		return err
	}

	return nil
}

func (d rootLayer) Create(data model.Layer) error { return nil }

func (d rootLayer) Override(m model.Layer) error { return nil }

func (d rootLayer) AddField(m model.Layer) error { return nil }
