package model

import (
	"os"
	"path/filepath"
)

// Config model to map configuration from a yaml file
type Config struct {
	ProjectPath  string     `yaml:"project_path"`
	ModuleName   ModuleName `yaml:"module_name"`
	Model        string     `yaml:"model"`
	Table        string     `yaml:"table"`
	TableComment string     `yaml:"table_comment"`
	Layers       []string   `yaml:"layers"`
	Fields       Fields     `yaml:"fields"`
	Architecture string
}

func (c *Config) SetInitPath(moduleName ModuleName) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	c.ProjectPath = filepath.Join(path, moduleName.GetProjectName())

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

func (c Config) IsProjectPathEmpty() bool {
	return c.ProjectPath == ""
}
