package model

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Layer struct {
	Model  string
	Table  string
	Fields Fields

	// ProjectPath indicates the root location of the project
	ProjectPath string

	// ModuleName is used to build the imports
	ModuleName string
}

// NewLayer returns a new Layer with module and table Field initialized
func NewLayer(conf Config) Layer {
	return Layer{
		ProjectPath: conf.ProjectPath,
		ModuleName:  conf.ModuleName,
		Model:       conf.Model,
		Table:       conf.Table,
		Fields:      conf.Fields,
	}
}

func (l *Layer) GetPath(layerName, filename string) string {
	return fmt.Sprintf("%s", filepath.Join(l.ProjectPath, layerName, strings.ToLower(l.Model), filename))
}
