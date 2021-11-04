package runner

import (
	"fmt"

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
	fmt.Println(getActionMessage(a, m))

	for k, layerUseCase := range r.layers {
		prefix := fmt.Sprintf("[%d/%d]", k+1, len(r.layers))

		if err := r.exec(a, m, layerUseCase); err != nil {
			layerUseCase.FailureMsg(prefix, err)
			continue
		}

		layerUseCase.SuccessfulMsg(prefix)
	}

	return nil
}

func (r runner) exec(a Action, m model.Layer, layerUseCase layer.UseCase) error {
	if a == Init {
		return layerUseCase.Init(m)
	}

	useCase, ok := layerUseCase.(layer.UseCasePackage)
	if !ok {
		return fmt.Errorf("layer does not implement the `%s` action", a)
	}

	switch a {
	case Override:
		return useCase.Override(m)
	case NewField:
		return useCase.AddField(m)
	case NewPackage:
		return useCase.Create(m)
	}

	return fmt.Errorf("action does not implemented by any layer")
}

func getActionMessage(action Action, m model.Layer) string {
	switch action {
	case Init:
		return fmt.Sprintf("Inicializando nuevo proyecto en %s", m.ProjectPath)
	case NewPackage:
		return fmt.Sprintf("Creando nuevo paquete en %s", m.ProjectPath)
	default:
		return ""
	}
}
