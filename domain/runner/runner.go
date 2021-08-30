package runner

import (
	"github.com/edteamlat/go-wizard/model"
)

type Action string

const (
	Override Action = "override"
	NewField Action = "new-field"
	Init     Action = "init"
)

type UseCase interface {
	GenerateLayers(a Action, m model.Layer) error
}
