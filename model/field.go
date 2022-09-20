package model

// Field model for every Field of a struct and table that want to be generated
type Field struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	IsNull bool   `yaml:"is_null"`
}

// Fields slice of Field
type Fields []Field

func (f Fields) HasID() bool {
	for _, field := range f {
		if field.Name == "id" {
			return true
		}
	}

	return false
}

func (f Fields) HasCreatedAt() bool {
	for _, field := range f {
		if field.Name == "created_at" {
			return true
		}
	}

	return false
}

