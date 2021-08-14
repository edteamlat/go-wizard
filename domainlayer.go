package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

const (
	domainTemplateName  = "domain.gotpl"
	useCaseTemplateName = "usecase.gotpl"
)

const (
	domainPath = "domain/%s/%s.go"
)

type domainLayer struct {
	template useCaseTemplate
	storage  storage
}

func newDomainLayer(template useCaseTemplate, storage storage) *domainLayer {
	return &domainLayer{template: template, storage: storage}
}

func (d domainLayer) create(data layer) error {
	if err := d.createDomainFile(data); err != nil {
		return fmt.Errorf("domainlayer: %w", err)
	}

	if err := d.createUseCaseFile(data); err != nil {
		return fmt.Errorf("domainlayer: %w", err)
	}

	return nil
}

func (d domainLayer) createDomainFile(data layer) error {
	domainFileBuf := bytes.Buffer{}
	if err := d.template.create(&domainFileBuf, domainTemplateName, data); err != nil {
		return err
	}

	packageName := strings.ToLower(data.ModelName)
	domainFilePath := fmt.Sprintf(domainPath, packageName, packageName)
	if err := d.storage.save(filepath.Join(data.ProjectPath, domainFilePath), domainFileBuf); err != nil {
		return err
	}

	return nil
}

func (d domainLayer) createUseCaseFile(data layer) error {
	useCaseFileBuf := bytes.Buffer{}
	if err := d.template.create(&useCaseFileBuf, useCaseTemplateName, data); err != nil {
		return err
	}

	packageName := strings.ToLower(data.ModelName)
	domainUseCaseFilePath := fmt.Sprintf(domainPath, packageName, "usecase")
	if err := d.storage.save(filepath.Join(data.ProjectPath, domainUseCaseFilePath), useCaseFileBuf); err != nil {
		return err
	}

	return nil
}

func (d domainLayer) override(m layer) error {
	panic("implement me")
}

func (d domainLayer) addField(m layer) error {
	panic("implement me")
}
