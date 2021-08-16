package filesystem

import (
	"bytes"
	"io/fs"
	"io/ioutil"
)

type fileSystem struct {
}

func NewFileSystem() *fileSystem {
	return &fileSystem{}
}

func (f fileSystem) Save(path string, data bytes.Buffer) error {
	ioutil.WriteFile(path, data.Bytes(), fs.ModePerm)

	return nil
}
