package layer

import (
	"bytes"
	"fmt"
	"io"

	"github.com/edteamlat/go-wizard/domain/edhex"
	"github.com/edteamlat/go-wizard/model"
)

// UseCase use case to generate a layer
type UseCase interface {
	Create(m model.Layer) error
	Override(m model.Layer) error
	AddField(m model.Layer) error
}

// UseCaseLayers slice of useCaseLayers
type UseCaseLayers []UseCase

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
			return nil, err
		}

		layers = append(layers, layer)
	}

	return layers, nil
}

// getLayer factory that obtains a new useCaseLayer
func getLayer(name string, template UseCaseTemplate, storage Storage) (UseCase, error) {
	switch name {
	case edhex.DomainLayerName:
		return edhex.NewDomainLayer(template, storage), nil
	default:
		return nil, fmt.Errorf("layer is not implemented")
	}
}
