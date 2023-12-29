package controllers

import (
	"fmt"
	"golang-go/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StorageController struct {
	storageService *services.StorageService
}

func NewStorageController(storageService *services.StorageService) *StorageController {
	return &StorageController{
		storageService: storageService,
	}

}

func (s StorageController) CheckFiles(server echo.Context) error {
	name := server.QueryParam("name")
	message := fmt.Sprintf("%s is checked", name)
	return response(server, http.StatusOK, 0, "success", message)
}
