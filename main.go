package main

import (
	"embed"
	"text/template"

	"github.com/edteamlat/go-wizard/domain/edhex"
	"github.com/edteamlat/go-wizard/domain/layer"
	"github.com/edteamlat/go-wizard/domain/runner"
	"github.com/edteamlat/go-wizard/infrastructure/filesystem"
	"github.com/edteamlat/go-wizard/infrastructure/texttemplate"
	"github.com/edteamlat/go-wizard/model"
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

	runnerUseCase, err := buildUseCaseRunner(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := runnerUseCase.GenerateLayers("", layerData); err != nil {
		log.Fatal(err)
	}
}

func buildUseCaseRunner(conf model.Config) (runner.UseCase, error) {
	layerUseCases, err := buildUseCaseLayers(conf)
	if err != nil {
		return nil, err
	}

	runnerUseCase := runner.NewRunner()
	runnerUseCase.AppendLayer(layerUseCases...)
	return runnerUseCase, nil
}

func buildUseCaseLayers(conf model.Config) (layer.UseCaseLayers, error) {
	fileSystemUseCase := filesystem.NewFileSystem()

	tpl := template.Must(template.New("").Funcs(edhex.GetTemplateFunctions()).ParseFS(templatesFS))
	templateUseCase := texttemplate.NewTemplate(tpl)

	return layer.GetUseCaseLayersFromConf(conf, templateUseCase, fileSystemUseCase)
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
