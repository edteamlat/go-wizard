package main

type Runner struct {
	layers UseCaseLayers
}

// NewRunner returns a new Runner
func NewRunner() *Runner {
	return &Runner{}
}

// AppendLayer adds a new UseCaseLayer to Runner.layers field
func (r *Runner) AppendLayer(layer ...UseCaseLayer) {
	r.layers = append(r.layers, layer...)
}

// Run runs the generation of every layer
func (r Runner) Run(m Layer) error {
	for _, layer := range r.layers {
		if err := layer.Generate(m); err != nil {
			return err
		}
	}

	return nil
}
