package main

import (
	"embed"
	"text/template"

	"github.com/edteamlat/go-wizard/domain"
	"github.com/edteamlat/go-wizard/filesystem"
	"github.com/edteamlat/go-wizard/model"
	"github.com/edteamlat/go-wizard/texttemplate"

	"github.com/labstack/gommon/log"
)

//go:embed templates
var templates embed.FS

func main() {
	// TODO: read path from flag
	conf, err := readConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fields, err := model.NewFieldsFromSliceString(conf.Fields)
	if err != nil {
		log.Fatal(err)
	}

	layerData := model.NewLayer(conf)
	layerData.SetFields(fields)

	fileSystemUseCase := filesystem.NewFileSystem()

	tpl := template.Must(template.New("").Funcs(domain.GetTemplateFunctions()).ParseFS(templates))
	templateUseCase := texttemplate.NewTemplate(tpl)

	layerUseCases, err := domain.GetUseCaseLayersFromConf(conf, templateUseCase, fileSystemUseCase)

	runner := domain.NewRunner()
	runner.AppendLayer(layerUseCases...)
	if err := runner.Run("", *layerData); err != nil {
		log.Fatal(err)
	}
}
