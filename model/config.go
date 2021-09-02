package model

import (
	"os"
	"path/filepath"
	"strings"
)

// Config model to map configuration from a yaml file
type Config struct {
	ProjectPath  string   `yaml:"project_path"`
	ModuleName   string   `yaml:"module_name"`
	Model        string   `yaml:"model"`
	Table        string   `yaml:"table"`
	TableComment string   `yaml:"table_comment"`
	Layers       []string `yaml:"layers"`
	Fields       Fields   `yaml:"fields"`
	Architecture string
}

func (c *Config) SetInitPath(moduleName string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	moduleSplit := strings.Split(moduleName, "/")
	projectName := moduleSplit[len(moduleSplit)-1]

	c.ProjectPath = filepath.Join(path, projectName)

	return nil
}

func (c *Config) AddDefaultInitLayers() {
	c.Layers = append(
		c.Layers,
		"root",
		"cmd",
		"domain",
		"handler_echo",
		"storage_postgres",
		"model",
		"sqlmigration_postgres",
	)
}
