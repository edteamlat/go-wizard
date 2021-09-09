package edhex

import (
	"fmt"
	"path/filepath"

	"github.com/edteamlat/go-wizard/model"
)

const handlerFolder = "infrastructure/handler"
const HandlerLayerName = "handler_echo"

var handlerInitActionTemplates = model.Templates{
	{
		Name:     "fields.gotpl",
		Filename: "fields.go",
		Path:     filepath.Join(handlerFolder, "request"),
	},
	{
		Name:     "parameter.gotpl",
		Filename: "parameter.go",
		Path:     filepath.Join(handlerFolder, "request"),
	},
	{
		Name:     "token.gotpl",
		Filename: "token.go",
		Path:     filepath.Join(handlerFolder, "request"),
	},
	{
		Name:     "message.gotpl",
		Filename: "message.go",
		Path:     filepath.Join(handlerFolder, "response"),
	},
	{
		Name:     "response.gotpl",
		Filename: "response.go",
		Path:     filepath.Join(handlerFolder, "response"),
	},
	{
		Name:     "router.gotpl",
		Filename: "router.go",
		Path:     handlerFolder,
	},
}

var handlerAddActionTemplates = model.Templates{
	{
		Name:        "handler.gotpl",
		Filename:    "handler.go",
		Path:        handlerFolder,
		WithPackage: true,
	},
	{
		Name:        "route.gotpl",
		Filename:    "route.go",
		Path:        handlerFolder,
		WithPackage: true,
	},
}

type handlerLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func NewHandlerLayer(template UseCaseTemplate, storage Storage) handlerLayer {
	return handlerLayer{template: template, storage: storage}
}

func (d handlerLayer) Init(data model.Layer) error {
	if err := d.storage.CreateDir(filepath.Join(data.ProjectPath, "infrastructure", "handler")); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	if err := bulkTemplates(d.template, d.storage, handlerInitActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	return nil
}

func (d handlerLayer) Create(data model.Layer) error {
	if err := bulkTemplates(d.template, d.storage, handlerAddActionTemplates, data); err != nil {
		return fmt.Errorf("edhex-domainlayer: %w", err)
	}

	return nil
}

func (d handlerLayer) Override(m model.Layer) error {
	return nil
}

func (d handlerLayer) AddField(m model.Layer) error {
	return nil
}

func (d handlerLayer) SuccessfulMsg(prefixCount string) {
	fmt.Printf("%s handler layer generated ✅\n", prefixCount)
}

func (d handlerLayer) FailureMsg(prefixCount string, err error) {
	fmt.Printf("%s handler layer failed 🚨, %s\n", prefixCount, err.Error())
}
