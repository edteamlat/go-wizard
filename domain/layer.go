package domain

import (
	"bytes"
	"fmt"
	"io"

	"github.com/edteamlat/go-wizard/model"
)

// available layers that can be generated
const (
	domainLayerOption      = "domain"
	postgresLayerOption    = "postgres"
	handlerEchoLayerOption = "handler:echo"
)

// UseCaseLayer use case to generate a layer
type UseCaseLayer interface {
	Create(m model.Layer) error
	Override(m model.Layer) error
	AddField(m model.Layer) error
}

// UseCaseLayers slice of useCaseLayers
type UseCaseLayers []UseCaseLayer

type UseCaseTemplate interface {
	Create(wr io.Writer, templateName string, data model.Layer) error
}

type Storage interface {
	Save(path string, data bytes.Buffer) error
}

// GetUseCaseLayersFromConf obtains all useCaseLayers that were specified on the Config file
func GetUseCaseLayersFromConf(conf model.Config, template UseCaseTemplate, storage Storage) (UseCaseLayers, error) {
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

// getLayer factory that obtains a new useCaseLayer
func getLayer(name string, template UseCaseTemplate, storage Storage) (UseCaseLayer, error) {
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
