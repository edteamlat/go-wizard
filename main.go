package main

import (
	"embed"
	"text/template"

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

	fields, err := newFieldsFromSliceString(conf.Fields)
	if err != nil {
		log.Fatal(err)
	}

	layerData := newLayer(conf)
	layerData.setFields(fields)

	fileSystemUseCase := newFileSystem()

	tpl := template.Must(template.New("").Funcs(getTemplateFunctions()).ParseFS(templates))
	templateUseCase := newTemplate(tpl)

	layerUseCases, err := getUseCaseLayersFromConf(conf, templateUseCase, fileSystemUseCase)

	runner := newRunner()
	runner.appendLayer(layerUseCases...)
	if err := runner.run("", *layerData); err != nil {
		log.Fatal(err)
	}
}
