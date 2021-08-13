package main

import "fmt"

// available layers that can be generated
const (
	domainLayer      = "domain"
	postgresLayer    = "postgres"
	handlerEchoLayer = "handler:echo"
)

// UseCaseLayer use case to generate a layer
type UseCaseLayer interface {
	Generate(m Layer) error
}

// UseCaseLayers slice of UseCaseLayers
type UseCaseLayers []UseCaseLayer

type Layer struct {
	ModelName string
	TableName string
	Fields    Fields

	// ProjectPath indicates the root location of the project
	ProjectPath string
}

// newLayer returns a new Layer with module and table field initialized
func newLayer(conf Config) *Layer {
	return &Layer{
		ModelName: conf.ModuleName,
		TableName: conf.TableName,
	}
}

// SetFields sets the fields that will be used to generate de Struct fields and table columns
func (l *Layer) SetFields(fields Fields) {
	l.Fields = fields
}

// getUseCaseLayersFromConf obtains all UseCaseLayers that were specified on the config file
func getUseCaseLayersFromConf(conf Config, template UseCaseTemplate, storage Storage) (UseCaseLayers, error) {
	layers := UseCaseLayers{}

	for _, layerName := range conf.Layers {
		layer, err := getLayer(layerName, template, storage)
		if err != nil {
			return layers, err
		}

		layers = append(layers, layer)
	}

	return layers, nil
}

// getLayer factory that obtains a new UseCaseLayer
func getLayer(name string, template UseCaseTemplate, storage Storage) (UseCaseLayer, error) {
	switch name {
	case domainLayer:
		return newDomainLayer(template, storage), nil
	case postgresLayer:
		return nil, nil
	case handlerEchoLayer:
		return nil, nil
	default:
		return nil, fmt.Errorf("layer is not implemented")
	}
}
