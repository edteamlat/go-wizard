package model

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Layer struct {
	Model      string
	Table      string
	Fields     Fields
	LayersPath map[string]string

	// ProjectPath indicates the root location of the project
	ProjectPath string
}

// NewLayer returns a new Layer with module and table Field initialized
func NewLayer(conf Config) *Layer {
	return &Layer{
		ProjectPath: conf.ProjectPath,
		Model:       conf.Model,
		Table:       conf.Table,
		LayersPath:  conf.Layers,
	}
}

// SetFields sets the fields that will be used to generate de Struct fields and table columns
func (l *Layer) SetFields(fields Fields) {
	l.Fields = fields
}

func (l *Layer) GetPath(layerName, filename string) string {
	return fmt.Sprintf("%s.go", filepath.Join(l.ProjectPath, layerName, strings.ToLower(l.Model), filename))
}
