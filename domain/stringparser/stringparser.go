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
		"parseToUpperCamelCase":                 parseToUpperCamelCase,
		"parseToUpper":                          parseToUpper,
		"parseToLower":                          parseToLower,
		"parseToKebabCase":                      parseToKebabCase,
		"parseToLowerCamelCase":                 parseToLowerCamelCase,
		"parseToSnakeCase":                      parseToSnakeCase,
		"parseToUpperSnakeCase":                 parseToUpperSnakeCase,
		"getFirstLetter":                        getFirstLetter,
		"increment":                             increment,
		"decrement":                             decrement,
		"parseToSqlType":                        parseToSqlType,
		"handleNull":                            handleNull,
		"handleNullOnScan":                      handleNullOnScan,
		"parseNullFieldsOnScan":                 parseNullFieldsOnScan,
		"printFieldsWithoutDefaults":            printFieldsWithoutDefaults,
		"printStorageFieldsWithoutDefaults":     printStorageFieldsWithoutDefaults,
		"printStorageFieldsWithoutDateDefaults": printStorageFieldsWithoutDateDefaults,
		"printMigrationFieldsWithoutDefaults":   printMigrationFieldsWithoutDefaults,
		"printStorageNullFields":                printStorageNullFields,
		"printStorageNullFieldsScan":            printStorageNullFieldsScan,
		"printStorageNullFieldsParse":           printStorageNullFieldsParse,
	}
}

func parseToUpperCamelCase(v string) string {
	return parseIdToID(strcase.UpperCamelCase(v))
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
	return parseIdToID(strcase.LowerCamelCase(v))
}

func parseToSnakeCase(v string) string {
	return strcase.SnakeCase(v)
}

func parseToUpperSnakeCase(v string) string {
	return strings.ToUpper(strcase.SnakeCase(v))
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

func parseIdToID(v string) string {
	return strings.ReplaceAll(v, "Id", "ID")
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
	case "json.RawMessage":
		return "JSON"
	default:
		return "UNKNOWN-TYPE"
	}
}

func handleNull(f model.Field) string {
	field := parseToUpperCamelCase(f.Name)

	if !f.IsNull {
		return fmt.Sprintf("m.%s", field)
	}

	switch f.Type {
	case "string":
		return fmt.Sprintf("nullhandler.StringToNull(m.%s)", field)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return fmt.Sprintf("nullhandler.IntToNull(int64(m.%s))", field)
	case "float32":
		return fmt.Sprintf("nullhandler.FloatToNull(float64(m.%s))", field)
	case "float64":
		return fmt.Sprintf("nullhandler.FloatToNull(m.%s)", field)
	case "time.Time":
		return fmt.Sprintf("nullhandler.TimeToNull(m.%s)", field)
	default:
		return fmt.Sprintf("invalid data type: %s", f.Type)
	}
}

func handleNullOnScan(f model.Field) string {
	if !f.IsNull {
		return ""
	}

	name := parseToLowerCamelCase(f.Name)

	switch f.Type {
	case "string":
		return fmt.Sprintf("%sNull := sql.NullString{}", name)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return fmt.Sprintf("%sNull := sql.NullInt64{}", name)
	case "float64", "float32":
		return fmt.Sprintf("%sNull := sql.NullFloat64{}", name)
	case "time.Time":
		return fmt.Sprintf("%sNull := sql.NullTime{}", name)
	case "bool":
		return fmt.Sprintf("%sNull := sql.NullBool{}", name)
	default:
		return fmt.Sprintf("invalid data type: %s", name)
	}
}

func parseNullFieldsOnScan(f model.Field) string {
	field := parseToUpperCamelCase(f.Name)
	if !f.IsNull {
		return ""
	}

	fieldNull := parseToLowerCamelCase(f.Name) + "Null"

	switch f.Type {
	case "string":
		return fmt.Sprintf("m.%s = %s.String", field, fieldNull)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return fmt.Sprintf("m.%s = %s(%s.Int64)", field, f.Type, fieldNull)
	case "float32":
		return fmt.Sprintf("m.%s = %s(%s.Float64)", field, f.Type, fieldNull)
	case "float64":
		return fmt.Sprintf("m.%s = %s.Float64", field, fieldNull)
	case "time.Time":
		return fmt.Sprintf("m.%s = %s.Time", field, fieldNull)
	case "bool":
		return fmt.Sprintf("m.%s = %s.Bool", field, fieldNull)
	default:
		return fmt.Sprintf("invalid data type: %s", fieldNull)
	}
}

func printFieldsWithoutDefaults(fields model.Fields) string {
	msg := ""
	for _, field := range fields {
		if isDefaultField(field.Name) {
			continue
		}

		msg += fmt.Sprintf("\"%s\",\n\t", field.Name)
	}

	return strings.TrimSpace(msg)
}

func printStorageFieldsWithoutDefaults(fields model.Fields) string {
	msg := ""
	for _, field := range fields {
		if isDefaultField(field.Name) {
			continue
		}

		msg += fmt.Sprintf("%s,\n\t", handleNull(field))
	}

	return strings.TrimSpace(msg)
}

func printStorageFieldsWithoutDateDefaults(fields model.Fields) string {
	msg := ""
	for _, field := range fields {
		if isDefaultDateField(field.Name) {
			continue
		}

		msg += fmt.Sprintf("%s,\n\t", handleNull(field))
	}

	return strings.TrimSpace(msg)
}

func printMigrationFieldsWithoutDefaults(fields model.Fields) string {
	msg := ""
	for _, field := range fields {
		if isDefaultField(field.Name) {
			continue
		}

		msg += fmt.Sprintf("%s %s%s,\n\t", field.Name, parseToSqlType(field.Type), parseNull(field.IsNull))
	}

	return strings.TrimSpace(msg)
}

func printStorageNullFields(fields model.Fields) string {
	msg := ""
	for _, field := range fields {
		if !field.IsNull {
			continue
		}

		msg += fmt.Sprintf("%s\n\t", handleNullOnScan(field))
	}

	return strings.TrimSpace(msg)
}

func printStorageNullFieldsScan(fields model.Fields) string {
	msg := ""
	for _, field := range fields {
		if field.IsNull {
			msg += fmt.Sprintf("&%sNull,\n\t", parseToLowerCamelCase(field.Name))
			continue
		}

		msg += fmt.Sprintf("&m.%s,\n\t", parseToUpperCamelCase(field.Name))
	}

	return strings.TrimSpace(msg)
}

func printStorageNullFieldsParse(fields model.Fields) string {
	msg := ""
	for _, field := range fields {
		if !field.IsNull {
			continue
		}

		msg += fmt.Sprintf("%s\n\t", parseNullFieldsOnScan(field))
	}

	return strings.TrimSpace(msg)
}

func isDefaultField(field string) bool {
	return field == "id" || field == "created_at" || field == "updated_at"
}

func isDefaultDateField(field string) bool {
	return field == "created_at" || field == "updated_at"
}

func parseNull(isNull bool) string {
	if isNull {
		return ""
	}

	return " NOT NULL"
}
