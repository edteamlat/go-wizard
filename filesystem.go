package main

import (
	"bytes"
	"io/fs"
	"io/ioutil"
)

type Storage interface {
	Save(path string, data bytes.Buffer) error
}

type FileSystem struct {
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (f FileSystem) Save(path string, data bytes.Buffer) error {
	ioutil.WriteFile(path, data.Bytes(), fs.ModePerm)

	return nil
}
