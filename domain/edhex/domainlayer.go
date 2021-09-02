package edhex

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/edteamlat/go-wizard/model"
)

const (
	domainTemplateName  = "domain.gotpl"
	useCaseTemplateName = "usecase.gotpl"
)

const DomainLayerName = "domain"

type domainLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewDomainLayer(template UseCaseTemplate, storage Storage) domainLayer {
	return domainLayer{template: template, storage: storage}
}

func (d domainLayer) Init(m model.Layer) error {
	return nil
}

func (d domainLayer) Create(data model.Layer) error {
	if err := d.createDomainFile(data); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	if err := d.createUseCaseFile(data); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	return nil
}

func (d domainLayer) createDomainFile(data model.Layer) error {
	domainFileBuf := bytes.Buffer{}
	if err := d.template.Create(&domainFileBuf, domainTemplateName, data); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.go", strings.ToLower(data.Model))
	if err := d.storage.Save(data.GetPath(DomainLayerName, filename, true), domainFileBuf); err != nil {
		return err
	}

	return nil
}

func (d domainLayer) createUseCaseFile(data model.Layer) error {
	useCaseFileBuf := bytes.Buffer{}
	if err := d.template.Create(&useCaseFileBuf, useCaseTemplateName, data); err != nil {
		return err
	}

	if err := d.storage.Save(data.GetPath(DomainLayerName, "usecase.go", true), useCaseFileBuf); err != nil {
		return err
	}

	return nil
}

func (d domainLayer) Override(m model.Layer) error {
	return nil
}

func (d domainLayer) AddField(m model.Layer) error {
	return nil
}
