package runner

import (
	"github.com/edteamlat/go-wizard/model"
)

type action string

const (
	override action = "override"
	newField action = "new-field"
)

type UseCase interface {
	GenerateLayers(a action, m model.Layer) error
}
