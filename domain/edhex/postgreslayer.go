package edhex

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/edteamlat/go-wizard/model"
)

const (
	postgresTemplateName  = "postgres.gotpl"
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

func (d postgresLayer) Create(data model.Layer) error {
	if err := d.createPostgres(data); err != nil {
		return fmt.Errorf("edhex-postgreslayer: %w", err)
	}

	return nil
}

func (d postgresLayer) createPostgres(data model.Layer) error {
	fileBuf := bytes.Buffer{}
	if err := d.template.Create(&fileBuf, postgresTemplateName, data); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.go", strings.ToLower(data.Model))
	if err := d.storage.Save(data.GetPath(postgresFolder, filename, false), fileBuf); err != nil {
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

