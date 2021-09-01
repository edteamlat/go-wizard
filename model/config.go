package model

// Config model to map configuration from a yaml file
type Config struct {
	ProjectPath  string   `yaml:"project_path"`
	ModuleName   string   `yaml:"module_name"`
	Model        string   `yaml:"model"`
	Table        string   `yaml:"table"`
	TableComment string   `yaml:"table_comment"`
	Layers       []string `yaml:"layers"`
	Fields       Fields   `yaml:"fields"`
}
