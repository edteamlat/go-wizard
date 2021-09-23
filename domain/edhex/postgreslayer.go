package edhex

import (
	"fmt"
	"path/filepath"

	"github.com/edteamlat/go-wizard/model"
)

const PostgresLayerName = "storage_postgres"

const postgresFolder = "infrastructure/postgres"

var postgresAddActionTemplates = model.Templates{
	{
		Name:        "postgres.gotpl",
		Filename:    "%s.go", // the name will be the name of the package
		Path:        postgresFolder,
		WithPackage: true,
	},
}

type postgresLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewPostgresLayer(template UseCaseTemplate, storage Storage) postgresLayer {
	return postgresLayer{template: template, storage: storage}
}

func (d postgresLayer) Init(data model.Layer) error {
	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "infrastructure", "postgres")); err != nil {
		return fmt.Errorf("edhex-postgreslayer: %w", err)
	}

	return nil
}

func (d postgresLayer) Create(data model.Layer) error {
	if err := bulkFromTemplates(d.template, d.storage, postgresAddActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-postgreslayer: %w", err)
	}

	return nil
}

func (d postgresLayer) Override(m model.Layer) error {
	return nil
}

func (d postgresLayer) AddField(m model.Layer) error {
	return nil
}

func (d postgresLayer) SuccessfulMsg(prefixCount string) {
	fmt.Printf("%s postgres layer generated ✅\n", prefixCount)
}

func (d postgresLayer) FailureMsg(prefixCount string, err error) {
	fmt.Printf("%s postgres layer failed 🚨, %s\n", prefixCount, err.Error())
}
