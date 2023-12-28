package routers

import (
	"golang-go/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(server *echo.Echo) {
	server.GET("/healthz", controllers.Healthcheck)
}

// func initv1(server *echo.Echo) {
// 	v1 := server.Group("/v1")
// }

// func initOperator(group *echo.Group) {
// }
