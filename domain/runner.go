package domain

import (
	"github.com/edteamlat/go-wizard/model"
)

type action string

const (
	override action = "override"
	newField action = "new-field"
)

type UseCaseRunner interface {
	GenerateLayers(a action, m model.Layer) error
}

type runner struct {
	layers UseCaseLayers
}

// NewRunner returns a new runner
func NewRunner() *runner {
	return &runner{}
}

// AppendLayer adds a new useCaseLayer to runner.layers field
func (r *runner) AppendLayer(layer ...UseCaseLayer) {
	r.layers = append(r.layers, layer...)
}

// GenerateLayers runs the generation of every layer
func (r runner) GenerateLayers(a action, m model.Layer) error {
	for _, layerUseCase := range r.layers {
		if err := r.exec(a, m, layerUseCase); err != nil {
			return err
		}
	}

	return nil
}

func (r runner) exec(a action, m model.Layer, layerUseCase UseCaseLayer) error {
	switch a {
	case override:
		return layerUseCase.Override(m)
	case newField:
		return layerUseCase.AddField(m)
	default:
		return layerUseCase.Create(m)
	}
}
