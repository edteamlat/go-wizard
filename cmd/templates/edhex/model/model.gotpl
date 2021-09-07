package model

import (
	"errors"
	"fmt"
	"reflect"
)

// Errors
var (
	ErrNilPointer        = errors.New("received model is null")
	ErrInvalidID         = errors.New("received ID is invalid")
	ErrModelNotFound     = errors.New("model not found")
	ErrParameterNotFound = errors.New("could not found the parameter")
	ErrRequestParamEmpty = errors.New("url parameter must not be empty")
)

// Errors SQL
var (
	ErrUnique     = errors.New("unique violation")
	ErrForeignKey = errors.New("foreign key violation")
	ErrNotNull    = errors.New("not Null violation")
)

// ValidateStructNil returns an error if the model is nil
func ValidateStructNil(i interface{}) error {
	// omit struct type
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		return nil
	}

	// Type: nil, Value: nil
	if i == nil {
		return ErrNilPointer
	}

	// Type: StructPointer, Value: nil
	// example: Type: *Configuration, Value: nil
	if reflect.ValueOf(i).IsNil() {
		return ErrNilPointer
	}

	// Type: StructPointer, Value: ZeroValue
	// example: Type: *CashBox, Value: &CashBox{}
	return nil
}

func errRequiredField(field string) error {
	e := NewError()
	e.SetError(fmt.Errorf("missing %s field", field))
	e.SetAPIMessage(fmt.Sprintf("¡Upps! no enviaste el campo %s", field))
	return e
}
