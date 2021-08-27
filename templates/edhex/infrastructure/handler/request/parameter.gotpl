package request

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// ExtractIDFromRequest extracts id of the c.Param
func ExtractIDFromURLParam(c echo.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}

// ExtractIDFromURLParamByName extracts id of the c.Param
func ExtractIDFromURLParamByName(name string, c echo.Context) (int, error) {
	return strconv.Atoi(c.Param(name))
}

// ExtractIDFromURLParamByName extracts id of the c.Param
func ExtractIDFromFormParamByName(name string, c echo.Context) (int, error) {
	paramValue := c.FormValue(name)
	if paramValue == "" {
		paramValue = "0"
	}

	return strconv.Atoi(paramValue)
}

// ExtractIDFromURLQueryParamByName extracts id of the c.QueryParam
func ExtractIDFromURLQueryParamByName(name string, c echo.Context) (int, error) {
	return strconv.Atoi(c.QueryParam(name))
}
