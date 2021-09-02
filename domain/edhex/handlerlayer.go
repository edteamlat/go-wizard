package edhex

import (
	"fmt"

	"github.com/edteamlat/go-wizard/model"
)

const (
	handlerTemplateName = "handler.gotpl"
	routeTemplateName   = "route.gotpl"
)

const HandlerLayerName = "handler_echo"

const handlerFolder = "infrastructure/handler"

type handlerLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewHandlerLayer(template UseCaseTemplate, storage Storage) handlerLayer {
	return handlerLayer{template: template, storage: storage}
}

func (d handlerLayer) Init(m model.Layer) error {
	return nil
}

func (d handlerLayer) Create(data model.Layer) error {
	if err := d.createHandler(data); err != nil {
		return fmt.Errorf("edhex-handlerlayer: %w", err)
	}

	if err := d.createRoute(data); err != nil {
		return fmt.Errorf("edhex-handlerlayer: %w", err)
	}

	return nil
}

func (d handlerLayer) createHandler(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  handlerTemplateName,
		Path:  data.GetPath(handlerFolder, "handler.go", true),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d handlerLayer) createRoute(data model.Layer) error {
	if err := createTemplate(d.template, d.storage, model.Template{
		Name:  routeTemplateName,
		Path:  data.GetPath(handlerFolder, "route.go", true),
		Layer: data,
	}); err != nil {
		return err
	}

	return nil
}

func (d handlerLayer) Override(m model.Layer) error {
	return nil
}

func (d handlerLayer) AddField(m model.Layer) error {
	return nil
}
