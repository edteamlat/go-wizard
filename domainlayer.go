package main

import (
	"bytes"
	"fmt"
)

const (
	domainTemplateName  = "domain.gotpl"
	useCaseTemplateName = "usecase.gotpl"
)

type DomainLayer struct {
	template UseCaseTemplate
	storage  Storage
}

func newDomainLayer(template UseCaseTemplate, storage Storage) *DomainLayer {
	return &DomainLayer{template: template, storage: storage}
}

func (d DomainLayer) Generate(data Layer) error {
	if err := d.generateDomainFile(data); err != nil {
		return fmt.Errorf("domainlayer: %w", err)
	}

	if err := d.generateUseCaseFile(data); err != nil {
		return fmt.Errorf("domainlayer: %w", err)
	}

	return nil
}

func (d DomainLayer) generateDomainFile(data interface{}) error {
	domainFileBuf := bytes.Buffer{}
	if err := d.template.Create(&domainFileBuf, domainTemplateName, &data); err != nil {
		return err
	}

	if err := d.storage.Save("", domainFileBuf); err != nil {
		return err
	}

	return nil
}

func (d DomainLayer) generateUseCaseFile(data interface{}) error {
	useCaseFileBuf := bytes.Buffer{}
	if err := d.template.Create(&useCaseFileBuf, useCaseTemplateName, &data); err != nil {
		return err
	}

	if err := d.storage.Save("", useCaseFileBuf); err != nil {
		return err
	}

	return nil
}
