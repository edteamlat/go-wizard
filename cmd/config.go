package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/edteamlat/go-wizard/domain/runner"
	"github.com/edteamlat/go-wizard/model"
)

func readConfig(filename string, action runner.Action) (model.Config, error) {
	if action == runner.Init {
		conf := model.Config{}
		conf.AddDefaultInitLayers()
		return conf, nil
	}

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return model.Config{}, fmt.Errorf("config: could not read file %s, %w", filename, err)
	}

	conf := model.Config{}
	if err := yaml.Unmarshal(fileBytes, &conf); err != nil {
		return conf, fmt.Errorf("config: could not unmarshal file, %w", err)
	}
	if !conf.IsProjectPathEmpty() {
		return conf, nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return conf, fmt.Errorf("config: could not get project path, %w", err)
	}

	conf.ProjectPath = dir

	return conf, nil
}
