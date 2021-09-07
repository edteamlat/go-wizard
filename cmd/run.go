package cmd

import (
	"embed"
	"text/template"

	"github.com/edteamlat/go-wizard/domain/layer"
	"github.com/edteamlat/go-wizard/domain/runner"
	"github.com/edteamlat/go-wizard/domain/stringparser"
	"github.com/edteamlat/go-wizard/infrastructure/filesystem"
	"github.com/edteamlat/go-wizard/infrastructure/texttemplate"
	"github.com/edteamlat/go-wizard/model"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

//go:embed templates
var templatesFS embed.FS

func run(cmd *cobra.Command, args []string, action runner.Action) {
	configPath := cmd.Flag(configPathFlag).Value.String()
	architecture := cmd.Flag(architectureFlag).Value.String()

	moduleName := ""
	if action == runner.Init {
		moduleName = cmd.Flag(moduleFlag).Value.String()
	}

	conf, err := readConfig(configPath, action)
	if err != nil {
		log.Fatal(err)
	}
	conf.Architecture = architecture
	if !isEmpty(moduleName) {
		if err := conf.SetInitPath(moduleName); err != nil {
			log.Fatal(err)
		}
		conf.ModuleName = moduleName
	}

	layerData := model.NewLayer(conf)

	runnerUseCase, err := buildUseCaseRunner(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := runnerUseCase.GenerateLayers(action, layerData); err != nil {
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
	fileSystemUseCase := filesystem.New()

	tpl, err := template.New("").Funcs(stringparser.GetTemplateFunctions()).ParseFS(
		templatesFS,
		"**/**/*.gotpl",
		"**/**/**/*.gotpl",
		"**/**/**/**/*.gotpl",
		"**/**/**/**/**/*.gotpl",
		"**/**/**/**/**/*.gotpl",
	)
	if err != nil {
		log.Fatal(err)
	}
	templateUseCase := texttemplate.NewTemplate(tpl)

	return layer.GetUseCaseLayersFromConf(conf, templateUseCase, fileSystemUseCase)
}

func isEmpty(s string) bool {
	return s == ""
}
