package stringparser

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/edteamlat/go-wizard/model"
	"github.com/stoewer/go-strcase"
)

func GetTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"parseToUpperCamelCase": parseToUpperCamelCase,
		"parseToUpper":          parseToUpper,
		"parseToLower":          parseToLower,
		"parseToKebabCase":      parseToKebabCase,
		"parseToLowerCamelCase": parseToLowerCamelCase,
		"getFirstLetter":        getFirstLetter,
		"increment":             increment,
		"decrement":             decrement,
		"parseToSqlType":        parseToSqlType,
		"handleNull":            handleNull,
		"handleNullOnScan":      handleNullOnScan,
		"parseNullFieldsOnScan": parseNullFieldsOnScan,
	}
}

func parseToUpperCamelCase(v string) string {
	return strcase.UpperCamelCase(v)
}

func parseToUpper(v string) string {
	return strings.ToUpper(v)
}

func parseToLower(v string) string {
	return strings.ToLower(v)
}

func parseToKebabCase(v string) string {
	return strcase.KebabCase(v)
}

func parseToLowerCamelCase(v string) string {
	return strcase.LowerCamelCase(v)
}

func getFirstLetter(v string) string {
	return strings.ToLower(string(v[0]))
}

func increment(v int) int {
	return v + 1
}

func decrement(v int) int {
	return v - 1
}

func parseToSqlType(v string) string {
	switch v {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "INTEGER"
	case "float64", "float32":
		return "NUMERIC(SIZE)"
	case "string":
		return "VARCHAR(SIZE)"
	case "bool":
		return "BOOLEAN"
	case "time.Time":
		return "TIMESTAMP"
	default:
		return "CHANGE-THIS-TYPE"
	}
}

func handleNull(f model.Field) string {
	field := strcase.UpperCamelCase(f.Name)

	if !f.IsNull {
		return fmt.Sprintf("m.%s", field)
	}

	switch f.Type {
	case "string":
		return fmt.Sprintf("sqlutil.StringToNull(m.%s)", field)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return fmt.Sprintf("sqlutil.IntToNull(int64(m.%s))", field)
	case "float64", "float32":
		return fmt.Sprintf("sqlutil.FloatToNull(float64(m.%s))", field)
	case "time.Time":
		return fmt.Sprintf("sqlutil.TimeToNull(m.%s)", field)
	default:
		return fmt.Sprintf("invalid data type: %s", f.Type)
	}
}

func handleNullOnScan(f model.Field) string {
	if !f.IsNull {
		return ""
	}

	switch f.Type {
	case "string":
		return fmt.Sprintf("%s := sql.NullString{}", f.Name)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return fmt.Sprintf("%s := sql.NullInt64{}", f.Name)
	case "float64", "float32":
		return fmt.Sprintf("%s := sql.NullFloat64{}", f.Name)
	case "time.Time":
		return fmt.Sprintf("%s := sql.NullTime{}", f.Name)
	case "bool":
		return fmt.Sprintf("%s := sql.NullBool{}", f.Name)
	default:
		return fmt.Sprintf("invalid data type: %s", f.Type)
	}
}

func parseNullFieldsOnScan(f model.Field) string {
	field := strcase.UpperCamelCase(f.Name)
	if !f.IsNull {
		return ""
	}

	switch f.Type {
	case "string":
		return fmt.Sprintf("m.%s = %s.String", field, f.Name)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return fmt.Sprintf("m.%s = %s(%s.Int64)", field, f.Type, f.Name)
	case "float64", "float32":
		return fmt.Sprintf("m.%s = %s(%s.Float64)", field, f.Type, f.Name)
	case "time.Time":
		return fmt.Sprintf("m.%s = %s.Time", field, f.Name)
	case "bool":
		return fmt.Sprintf("m.%s = %s.Bool", field, f.Name)
	default:
		return fmt.Sprintf("invalid data type: %s", f.Type)
	}
}
