package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/labstack/gommon/log"
)

// config model to map configuration from a yaml file
type config struct {
	ProjectPath string   `yaml:"project_path"`
	ModuleName  string   `yaml:"module_name"`
	TableName   string   `yaml:"table_name"`
	Fields      []string `yaml:"fields"`
	Layers      []string `yaml:"layers"`
}

func readConfig(filename string) (config, error) {
	log.Infof("Loading configuration file from %s...", filename)

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return config{}, fmt.Errorf("config: %w", err)
	}

	conf := config{}
	if err := yaml.Unmarshal(fileBytes, &conf); err != nil {
		return conf, fmt.Errorf("config: %w", err)
	}

	log.Info("Configuration file has been loaded.")

	return conf, nil
}
