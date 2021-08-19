package layer

import (
	"fmt"

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

// GetUseCaseLayersFromConf obtains all useCaseLayers that were specified on the Config file
func GetUseCaseLayersFromConf(conf model.Config, template edhex.UseCaseTemplate, storage edhex.Storage) (UseCaseLayers, error) {
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
func getLayer(name string, template edhex.UseCaseTemplate, storage edhex.Storage) (UseCase, error) {
	switch name {
	case edhex.DomainLayerName:
		return edhex.NewDomainLayer(template, storage), nil
	default:
		return nil, fmt.Errorf("layer is not implemented")
	}
}
