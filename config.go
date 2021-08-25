package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/edteamlat/go-wizard/model"
	"gopkg.in/yaml.v3"
)

func readConfig(filename string) (model.Config, error) {
	log.Printf("Loading configuration file from %s...", filename)

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return model.Config{}, fmt.Errorf("config: could not read file %s, %w", filename, err)
	}

	conf := model.Config{}
	if err := yaml.Unmarshal(fileBytes, &conf); err != nil {
		return conf, fmt.Errorf("config: could not unmarshal file, %w", err)
	}

	log.Println("Configuration file has been loaded.")

	return conf, nil
}
