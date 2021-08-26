package filesystem

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

const goExt = ".go"

type FileSystem struct{}

func New() FileSystem {
	return FileSystem{}
}

func (f FileSystem) Save(path string, buffer bytes.Buffer) error {
	dir := filepath.Dir(path)
	if err := createDir(dir); err != nil {
		return fmt.Errorf("filesystem: path: %s, %w, %v", dir, err, reflect.TypeOf(err))
	}

	if err := ioutil.WriteFile(path, getFormattedFile(path, buffer), os.ModePerm); err != nil {
		return fmt.Errorf("filesystem: could not create file %s, %w", path, err)
	}

	return nil
}

func createDir(dir string) error {
	d, err := os.Stat(dir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	// if dir already exists we dont have to create it
	if d != nil && d.IsDir() {
		return nil
	}

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func getFormattedFile(path string, buffer bytes.Buffer) []byte {
	if !isGolangFile(path) {
		return buffer.Bytes()
	}

	data, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Printf("filesystem: could not format file %s, err %v\n", path, err)
		return buffer.Bytes()
	}

	return data
}

func isGolangFile(path string) bool {
	return filepath.Ext(path) == goExt
}
