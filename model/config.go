package model

// Config model to map configuration from a yaml file
type Config struct {
	ProjectPath string `yaml:"project_path"`
	Model  string   `yaml:"module"`
	Table  string   `yaml:"table_name"`
	Fields []string `yaml:"fields"`
	Layers      []string `yaml:"layers"`
}
