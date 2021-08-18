package model

import (
	"errors"
	"strings"
)

var errInvalidField = errors.New("invalid len of Field")

const fieldSeparator = ":"

const typeFieldIndex = 0

const notNull = "NOT NULL"

// Field model for every Field of a struct and table that want to be generated
type Field struct {
	Name    string
	Type    string
	NotNull string
}

// Fields slice of Field
type Fields []Field

// NewFieldsFromMap builds a new Fields from fields in string format
func NewFieldsFromMap(fieldsMap map[string]string) (Fields, error) {
	fields := Fields{}

	for fieldName, fieldType := range fieldsMap {
		splitField := strings.Split(fieldType, fieldSeparator)
		if !isValidLen(splitField) {
			return fields, errInvalidField
		}

		field := Field{
			Name: fieldName,
			Type: splitField[typeFieldIndex],
		}

		if isNameTypeAndNotNull(splitField) {
			field.NotNull = notNull
		}

		fields = append(fields, field)
	}

	return fields, nil
}

func isValidLen(splitField []string) bool {
	fieldsLen := len(splitField)
	return fieldsLen >= 1 && fieldsLen <= 2
}

func isNameTypeAndNotNull(splitField []string) bool {
	return len(splitField) == 2
}
