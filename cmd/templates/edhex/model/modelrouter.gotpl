package model

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

// RouterSpecification is used to inject the router dependencies
type RouterSpecification struct {
	Config       Configuration
	Api          *echo.Echo
	Logger       Logger
	DB           *sql.DB
	RemoteConfig RemoteConfig
}
