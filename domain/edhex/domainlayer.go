package edhex

import (
	"fmt"
	"path/filepath"

	"github.com/edteamlat/go-wizard/model"
)

const DomainLayerName = "domain"

const domainFolder = "domain"

var domainAddActionTemplates = model.Templates{
	{
		Name:        "domain.gotpl",
		Filename:    "%s.go", // the name will be the name of the package
		Path:        domainFolder,
		WithPackage: true,
	},
	{
		Name:        "usecase.gotpl",
		Filename:    "usecase.go",
		Path:        domainFolder,
		WithPackage: true,
	},
}

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
	if err := bulkTemplates(d.template, d.storage, domainAddActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	return nil
}

func (d domainLayer) Override(m model.Layer) error {
	return nil
}

func (d domainLayer) AddField(m model.Layer) error {
	return nil
}

func (d domainLayer) SuccessfulMsg(prefixCount string) {
	fmt.Printf("%s domain layer generated ✅\n", prefixCount)
}

func (d domainLayer) FailureMsg(prefixCount string, err error) {
	fmt.Printf("%s domain layer failed 🚨, %s\n", prefixCount, err.Error())
}
