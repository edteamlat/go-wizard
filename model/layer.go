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

	TimeType TimeEnum `yaml:"time_type"`
	IDType   IDEnum   `yaml:"id_type"`
}

// NewLayer returns a new Layer with module and table Field initialized
func NewLayer(conf Config) Layer {

	// adds the default fields
	fields := Fields{{
		Name: "id",
		Type: string(conf.IDType),
	}}
	fields = append(fields, conf.Fields...)

	fields = append(fields, Fields{{
		Name: "created_at",
		Type: string(conf.TimeType),
	}, {
		Name:   "updated_at",
		Type:   string(conf.TimeType),
		IsNull: true,
	}}...)

	return Layer{
		ProjectPath:  conf.ProjectPath,
		ModuleName:   conf.ModuleName,
		Model:        conf.Model,
		Table:        conf.Table,
		TableComment: conf.TableComment,
		Fields:       fields,
		TimeType:     conf.TimeType,
		IDType:       conf.IDType,
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
