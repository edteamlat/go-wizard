package edhex

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/edteamlat/go-wizard/model"
)

const (
	postgresTemplateName = "postgres.gotpl"
)

const PostgresLayerName = "storage_postgres"

const postgresFolder = "infrastructure/postgres"

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
	if err := d.createPostgres(data); err != nil {
		return fmt.Errorf("edhex-postgreslayer: %w", err)
	}

	return nil
}

func (d postgresLayer) createPostgres(data model.Layer) error {
	filename := fmt.Sprintf("%s.go", strings.ToLower(data.Model))

	if err := createFromTemplate(d.template, d.storage, model.Template{
		Name:  postgresTemplateName,
		Path:  data.GetPath(postgresFolder, filename, true),
		Layer: data,
	}); err != nil {
		return err
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
