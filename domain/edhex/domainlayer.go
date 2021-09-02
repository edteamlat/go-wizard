package edhex

import (
	"fmt"
	"path/filepath"
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

func (d domainLayer) Init(data model.Layer) error {
	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "domain")); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}


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
	filename := fmt.Sprintf("%s.go", strings.ToLower(data.Model))

	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  domainTemplateName,
		Path:  data.GetPath(DomainLayerName, filename, true),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d domainLayer) createUseCaseFile(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  useCaseTemplateName,
		Path:  data.GetPath(DomainLayerName, "usecase.go", true),
		Layer: data,
	}); err != nil {
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
