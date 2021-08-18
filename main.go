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
var templatesFS embed.FS

func main() {
	// TODO: read path from flag
	conf, err := readConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	layerData, err := buildLayerModel(conf)
	if err != nil {
		log.Fatal(err)
	}

	runner, err := buildUseCaseRunner(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := runner.GenerateLayers("", layerData); err != nil {
		log.Fatal(err)
	}
}

func buildUseCaseRunner(conf model.Config) (domain.UseCaseRunner, error) {
	layerUseCases, err := buildUseCaseLayers(conf)
	if err != nil {
		return nil, err
	}

	runner := domain.NewRunner()
	runner.AppendLayer(layerUseCases...)
	return runner, nil
}

func buildUseCaseLayers(conf model.Config) (domain.UseCaseLayers, error) {
	fileSystemUseCase := filesystem.NewFileSystem()

	tpl := template.Must(template.New("").Funcs(domain.GetTemplateFunctions()).ParseFS(templatesFS))
	templateUseCase := texttemplate.NewTemplate(tpl)

	return domain.GetUseCaseLayersFromConf(conf, templateUseCase, fileSystemUseCase)
}

func buildLayerModel(conf model.Config) (model.Layer, error) {
	fields, err := model.NewFieldsFromMap(conf.Fields)
	if err != nil {
		return model.Layer{}, err
	}
	layerData := model.NewLayer(conf)
	layerData.SetFields(fields)

	return *layerData, nil
}
