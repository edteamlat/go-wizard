package domain

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/edteamlat/go-wizard/model"
)

const (
	domainTemplateName  = "domain.gotpl"
	useCaseTemplateName = "domain.gotpl"
)

const (
	domainPath = "domain/%s/%s.go"
)

type domainLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func newDomainLayer(template UseCaseTemplate, storage Storage) *domainLayer {
	return &domainLayer{template: template, storage: storage}
}

func (d domainLayer) Create(data model.Layer) error {
	if err := d.createDomainFile(data); err != nil {
		return fmt.Errorf("domainlayer: %w", err)
	}

	if err := d.createUseCaseFile(data); err != nil {
		return fmt.Errorf("domainlayer: %w", err)
	}

	return nil
}

func (d domainLayer) createDomainFile(data model.Layer) error {
	domainFileBuf := bytes.Buffer{}
	if err := d.template.Create(&domainFileBuf, domainTemplateName, data); err != nil {
		return err
	}

	packageName := strings.ToLower(data.Model)
	domainFilePath := fmt.Sprintf(domainPath, packageName, packageName)
	if err := d.storage.Save(filepath.Join(data.ProjectPath, domainFilePath), domainFileBuf); err != nil {
		return err
	}

	return nil
}

func (d domainLayer) createUseCaseFile(data model.Layer) error {
	useCaseFileBuf := bytes.Buffer{}
	if err := d.template.Create(&useCaseFileBuf, useCaseTemplateName, data); err != nil {
		return err
	}

	packageName := strings.ToLower(data.Model)
	domainUseCaseFilePath := fmt.Sprintf(domainPath, packageName, "domain")
	if err := d.storage.Save(filepath.Join(data.ProjectPath, domainUseCaseFilePath), useCaseFileBuf); err != nil {
		return err
	}

	return nil
}

func (d domainLayer) Override(m model.Layer) error {
	panic("implement me")
}

func (d domainLayer) AddField(m model.Layer) error {
	panic("implement me")
}
