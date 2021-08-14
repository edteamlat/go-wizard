package main

import (
	"bytes"
	"io/fs"
	"io/ioutil"
)

type storage interface {
	save(path string, data bytes.Buffer) error
}

type fileSystem struct {
}

func newFileSystem() *fileSystem {
	return &fileSystem{}
}

func (f fileSystem) save(path string, data bytes.Buffer) error {
	ioutil.WriteFile(path, data.Bytes(), fs.ModePerm)

	return nil
}
