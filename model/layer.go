package model

import (
	"fmt"
	"path/filepath"
	"strings"
)

type ModuleName string

func (m ModuleName) GetProjectName() string {
	moduleSplit := strings.Split(string(m), "/")
	return moduleSplit[len(moduleSplit)-1]
}

type Layer struct {
	Model        string
	Table        string
	TableComment string
	Fields       Fields

	// ProjectPath indicates the root location of the project
	ProjectPath string

	// ModuleName is used to build the imports
	ModuleName ModuleName
}

// NewLayer returns a new Layer with module and table Field initialized
func NewLayer(conf Config) Layer {

	// adds the default fields
	fields := Fields{{
		Name: "id",
		Type: "uint",
	}}
	fields = append(fields, conf.Fields...)
	fields = append(fields, Fields{{
		Name: "created_at",
		Type: "time.Time",
	}, {
		Name:   "updated_at",
		Type:   "time.Time",
		IsNull: true,
	}}...)

	return Layer{
		ProjectPath:  conf.ProjectPath,
		ModuleName:   conf.ModuleName,
		Model:        conf.Model,
		Table:        conf.Table,
		TableComment: conf.TableComment,
		Fields:       fields,
	}
}

func (l *Layer) GetPath(layerName, filename string, withPackage bool) string {
	if strings.HasPrefix(filename, "%s.") {
		packageName := strings.ToLower(l.Model)
		filename = fmt.Sprintf(filename, packageName)
	}

	if withPackage {
		return filepath.Join(l.ProjectPath, layerName, strings.ToLower(l.Model), filename)
	}

	return filepath.Join(l.ProjectPath, layerName, filename)
}
