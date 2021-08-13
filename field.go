package main

import (
	"errors"
	"strings"
)

var ErrInvalidField = errors.New("invalid len of field")

const fieldSeparator = ":"

const (
	nameFieldIndex = 0
	typeFieldIndex = 1
)

// Field model for every field of a struct and table that want to be generated
type Field struct {
	Name    string
	Type    string
	NotNull string
}

// Fields slice of Field
type Fields []Field

// newFieldsFromSliceString builds a new Fields from fields in string format
func newFieldsFromSliceString(fieldsStr []string) (Fields, error) {
	fields := Fields{}

	for _, fieldStr := range fieldsStr {
		splitField := strings.Split(fieldStr, fieldSeparator)
		if !isValidLen(splitField) {
			return fields, ErrInvalidField
		}

		field := Field{
			Name: splitField[nameFieldIndex],
			Type: splitField[typeFieldIndex],
		}

		if isNameTypeAndNotNull(splitField) {
			field.NotNull = "NOT NULL"
		}

		fields = append(fields, field)
	}

	return fields, nil
}

func isValidLen(splitField []string) bool {
	fieldsLen := len(splitField)
	return fieldsLen >= 2 && fieldsLen <= 3
}

func isNameTypeAndNotNull(splitField []string) bool {
	return len(splitField) == 2
}
