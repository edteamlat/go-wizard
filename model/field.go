package model

// Field model for every Field of a struct and table that want to be generated
type Field struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	IsNull    bool   `yaml:"is_null"`
	FieldSize int    `yaml:"field_size"`
}

// Fields slice of Field
type Fields []Field
