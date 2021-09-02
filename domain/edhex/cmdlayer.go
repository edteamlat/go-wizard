package edhex

import (
	"fmt"

	"github.com/edteamlat/go-wizard/model"
)

const (
	cmdConfigTemplateName       = "config.gotpl"
	cmdDatabaseTemplateName     = "database.gotpl"
	cmdEchoTemplateName         = "echo.gotpl"
	cmdLoggerTemplateName       = "logger.gotpl"
	cmdMainTemplateName         = "main.gotpl"
	cmdRemoteConfigTemplateName = "remoteconfig.gotpl"
)

const CMDLayerName = "cmd"

const packageName = "edhex-cmdlayer"

type cmdLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewCMDLayer(template UseCaseTemplate, storage Storage) cmdLayer {
	return cmdLayer{template: template, storage: storage}
}

func (d cmdLayer) Init(m model.Layer) error {
	if err := d.createConfig(m); err != nil {
		return fmt.Errorf("%s: %w", packageName, err)
	}

	if err := d.createDatabase(m); err != nil {
		return fmt.Errorf("%s: %w", packageName, err)
	}

	if err := d.createDatabase(m); err != nil {
		return fmt.Errorf("%s: %w", packageName, err)
	}

	if err := d.createEcho(m); err != nil {
		return fmt.Errorf("%s: %w", packageName, err)
	}

	if err := d.createLogger(m); err != nil {
		return fmt.Errorf("%s: %w", packageName, err)
	}

	if err := d.createMain(m); err != nil {
		return fmt.Errorf("%s: %w", packageName, err)
	}

	if err := d.createRemoteConfig(m); err != nil {
		return fmt.Errorf("%s: %w", packageName, err)
	}

	return nil
}

func (d cmdLayer) createConfig(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  cmdConfigTemplateName,
		Path:  data.GetPath(CMDLayerName, "config.go", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d cmdLayer) createDatabase(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  cmdDatabaseTemplateName,
		Path:  data.GetPath(CMDLayerName, "database.go", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d cmdLayer) createEcho(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  cmdEchoTemplateName,
		Path:  data.GetPath(CMDLayerName, "http.go", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d cmdLayer) createLogger(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  cmdLoggerTemplateName,
		Path:  data.GetPath(CMDLayerName, "logger.go", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d cmdLayer) createMain(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  cmdMainTemplateName,
		Path:  data.GetPath(CMDLayerName, "main.go", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d cmdLayer) createRemoteConfig(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  cmdRemoteConfigTemplateName,
		Path:  data.GetPath(CMDLayerName, "remoteconfig.go", false),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d cmdLayer) Create(data model.Layer) error { return nil }

func (d cmdLayer) Override(m model.Layer) error { return nil }

func (d cmdLayer) AddField(m model.Layer) error { return nil }
