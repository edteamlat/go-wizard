package main

type action string

const (
	override = "override"
	newField = "new-field"
)

type runner struct {
	layers useCaseLayers
}

// newRunner returns a new runner
func newRunner() *runner {
	return &runner{}
}

// appendLayer adds a new useCaseLayer to runner.layers field
func (r *runner) appendLayer(layer ...useCaseLayer) {
	r.layers = append(r.layers, layer...)
}

// run the generation of every layer
func (r runner) run(a action, m layer) error {
	for _, layerUseCase := range r.layers {
		if err := r.exec(a, m, layerUseCase); err != nil {
			return err
		}
	}

	return nil
}

func (r runner) exec(a action, m layer, layerUseCase useCaseLayer) error {
	switch a {
	case override:
		return layerUseCase.override(m)
	case newField:
		return layerUseCase.addField(m)
	default:
		return layerUseCase.create(m)
	}
}
