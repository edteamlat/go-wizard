package runner

import (
	"github.com/edteamlat/go-wizard/domain/layer"
	"github.com/edteamlat/go-wizard/model"
)

type runner struct {
	layers layer.UseCaseLayers
}

// NewRunner returns a new runner
func NewRunner() *runner {
	return &runner{}
}

// AppendLayer adds a new useCaseLayer to runner.layers field
func (r *runner) AppendLayer(layer ...layer.UseCase) {
	r.layers = append(r.layers, layer...)
}

// GenerateLayers runs the generation of every layer
func (r runner) GenerateLayers(a Action, m model.Layer) error {
	for _, layerUseCase := range r.layers {
		if err := r.exec(a, m, layerUseCase); err != nil {
			return err
		}
	}

	return nil
}

func (r runner) exec(a Action, m model.Layer, layerUseCase layer.UseCase) error {
	switch a {
	case Override:
		return layerUseCase.Override(m)
	case NewField:
		return layerUseCase.AddField(m)
	case Init:
		return layerUseCase.Init(m)
	default:
		return layerUseCase.Create(m)
	}
}
