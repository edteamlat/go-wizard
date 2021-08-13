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
	layerData.SetFields(fields)

	fileSystemUseCase := NewFileSystem()

	tpl := template.Must(template.New("").ParseFS(templates))
	templateUseCase := NewTemplate(tpl)

	layerUseCases, err := getUseCaseLayersFromConf(conf, templateUseCase, fileSystemUseCase)

	runner := NewRunner()
	runner.AppendLayer(layerUseCases...)
	if err := runner.Run(*layerData); err != nil {
		log.Fatal(err)
	}
}
