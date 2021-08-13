package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/labstack/gommon/log"
)

// Config model to map configuration from a yaml file
type Config struct {
	ModuleName string   `yaml:"module_name"`
	TableName  string   `yaml:"table_name"`
	Fields     []string `yaml:"fields"`
	Layers     []string `yaml:"layers"`

	// IsOverride indicates if we want to regenerate the layers
	// is set by flags
	IsOverride bool
}

func readConfig(filename string) (Config, error) {
	log.Infof("Loading configuration file from %s...", filename)

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("config: %w", err)
	}

	conf := Config{}
	if err := yaml.Unmarshal(fileBytes, &conf); err != nil {
		return conf, fmt.Errorf("config: %w", err)
	}

	log.Info("Configuration file has been loaded.")

	return conf, nil
}
