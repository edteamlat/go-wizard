package edhex

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/edteamlat/go-wizard/model"
)

const (
	sqlMigrationTemplateName = "sqlmigration.gotpl"
)

const SQLMigrationLayerName = "sqlmigration_postgres"

const sqlMigrationFolder = "sqlmigration"

type sqlMigrationLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewSQLMigrationLayer(template UseCaseTemplate, storage Storage) sqlMigrationLayer {
	return sqlMigrationLayer{template: template, storage: storage}
}

func (d sqlMigrationLayer) Init(data model.Layer) error {
	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "sqlmigration")); err != nil {
		return fmt.Errorf("edhex-sqlmigration: %w", err)
	}

	return nil
}

func (d sqlMigrationLayer) Create(data model.Layer) error {
	if err := d.createSQLMigration(data); err != nil {
		return fmt.Errorf("edhex-sqlmigration: %w", err)
	}

	return nil
}

func (d sqlMigrationLayer) createSQLMigration(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  sqlMigrationTemplateName,
		Path:  data.GetPath(sqlMigrationFolder, getFilename(data.Table), false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func getFilename(table string) string {
	now := time.Now()
	return now.Format("20060102") + "_" + now.Format("150405") + "_create_" + table + "_table.sql"
}

func (d sqlMigrationLayer) Override(m model.Layer) error {
	return nil
}

func (d sqlMigrationLayer) AddField(m model.Layer) error {
	return nil
}
