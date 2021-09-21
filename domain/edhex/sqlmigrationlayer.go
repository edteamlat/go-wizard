package edhex

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/edteamlat/go-wizard/model"
)

const SQLMigrationLayerName = "sqlmigration_postgres"

const sqlMigrationFolder = "sqlmigration"

var sqlmigrationAddActionTemplates = model.Templates{
	{
		Name: "sqlmigration.gotpl",
		Path: sqlMigrationFolder,
	},
}

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
	d.setFilenameToTemplates(sqlmigrationAddActionTemplates, data.Table)

	if err := bulkFromTemplates(d.template, d.storage, sqlmigrationAddActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-sqlmigration: %w", err)
	}

	return nil
}

func (d sqlMigrationLayer) setFilenameToTemplates(templates model.Templates, table string) {
	for k := range templates {
		templates[k].Filename = getFilename(table)
	}
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

func (d sqlMigrationLayer) SuccessfulMsg(prefixCount string) {
	fmt.Printf("%s sql-migration layer generated ✅\n", prefixCount)
}

func (d sqlMigrationLayer) FailureMsg(prefixCount string, err error) {
	fmt.Printf("%s sql-migration layer failed 🚨, %s\n", prefixCount, err.Error())
}
