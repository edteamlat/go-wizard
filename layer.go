package main

import "fmt"

// available layers that can be generated
const (
	domainLayerOption      = "domain"
	postgresLayerOption    = "postgres"
	handlerEchoLayerOption = "handler:echo"
)

// useCaseLayer use case to generate a layer
type useCaseLayer interface {
	create(m layer) error
	override(m layer) error
	addField(m layer) error
}

// useCaseLayers slice of useCaseLayers
type useCaseLayers []useCaseLayer

type layer struct {
	ModelName string
	TableName string
	Fields    Fields

	// ProjectPath indicates the root location of the project
	ProjectPath string
}

// newLayer returns a new layer with module and table field initialized
func newLayer(conf config) *layer {
	return &layer{
		ProjectPath: conf.ProjectPath,
		ModelName:   conf.ModuleName,
		TableName:   conf.TableName,
	}
}

// setFields sets the fields that will be used to generate de Struct fields and table columns
func (l *layer) setFields(fields Fields) {
	l.Fields = fields
}

// getUseCaseLayersFromConf obtains all useCaseLayers that were specified on the config file
func getUseCaseLayersFromConf(conf config, template useCaseTemplate, storage storage) (useCaseLayers, error) {
	layers := useCaseLayers{}

	for _, layerName := range conf.Layers {
		layer, err := getLayer(layerName, template, storage)
		if err != nil {
			return layers, err
		}

		layers = append(layers, layer)
	}

	return layers, nil
}

// getLayer factory that obtains a new useCaseLayer
func getLayer(name string, template useCaseTemplate, storage storage) (useCaseLayer, error) {
	switch name {
	case domainLayerOption:
		return newDomainLayer(template, storage), nil
	case postgresLayerOption:
		return nil, nil
	case handlerEchoLayerOption:
		return nil, nil
	default:
		return nil, fmt.Errorf("layer is not implemented")
	}
}
