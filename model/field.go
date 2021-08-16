package model

import (
	"errors"
	"strings"
)

var errInvalidField = errors.New("invalid len of Field")

const fieldSeparator = ":"

const (
	nameFieldIndex = 0
	typeFieldIndex = 1
)

// Field model for every Field of a struct and table that want to be generated
type Field struct {
	Name    string
	Type    string
	NotNull string
}

// Fields slice of Field
type Fields []Field

// NewFieldsFromSliceString builds a new Fields from fields in string format
func NewFieldsFromSliceString(fieldsStr []string) (Fields, error) {
	fields := Fields{}

	for _, fieldStr := range fieldsStr {
		splitField := strings.Split(fieldStr, fieldSeparator)
		if !isValidLen(splitField) {
			return fields, errInvalidField
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
