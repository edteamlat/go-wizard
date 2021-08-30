package model

// Config model to map configuration from a yaml file
type Config struct {
	ProjectPath  string   `yaml:"project_path"`
	Model        string   `yaml:"module"`
	Table        string   `yaml:"table"`
	Layers       []string `yaml:"layers"`
	Fields       Fields   `yaml:"fields"`
	Architecture string
}
