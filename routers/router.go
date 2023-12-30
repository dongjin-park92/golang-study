package routers

import (
	"context"
	"golang-go/controllers"
	"golang-go/gateways"
	"golang-go/services"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ControllerFactory struct {
	s3client gateways.AWSS3Gateway
}

func NewControllerFactory() *ControllerFactory {
	return &ControllerFactory{
		s3client: *gateways.NewS3Client(context.Background()),
	}
}

func (f ControllerFactory) CreateStorageController() controllers.StorageController {
	storageService := services.NewStorageService(&f.s3client)
	return *controllers.NewStorageController(storageService)
}

func InitRoute(server *echo.Echo) {
	server.GET("/healthz", controllers.Healthcheck)
	initv1(server, NewControllerFactory())
	// Echo framwork에서 metrics 노출하기 위함
	server.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

}

func initv1(server *echo.Echo, factory *ControllerFactory) {
	v1 := server.Group("/v1")
	storageController := factory.CreateStorageController()
	v1.GET("/storage/:bucket/files", storageController.CheckFiles)
}
