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
	bucket := server.Param("bucket")
	err := s.storageService.GetObjectExist(server.Request().Context(), bucket)
	if err != nil {
		response(server, http.StatusInternalServerError, 600, "failure", err)
	}
	message := fmt.Sprintf("%s is checked", bucket)
	return response(server, http.StatusOK, 0, "success", message)
}
