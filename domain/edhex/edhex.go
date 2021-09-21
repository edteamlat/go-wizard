package edhex

import (
	"bytes"
	"io"

	"github.com/edteamlat/go-wizard/model"
)

type UseCaseTemplate interface {
	Create(wr io.Writer, templateName string, data model.Layer) error
}

type Storage interface {
	CreateDir(dir string) error
	Save(path string, data bytes.Buffer) error
}
