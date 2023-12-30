package loggers

import "github.com/labstack/echo/v4"

var Loggers echo.Logger

func InitLogger(server *echo.Echo) {
	Loggers = server.Logger
}
