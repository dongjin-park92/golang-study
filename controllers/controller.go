package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 정상일 경우 errorCode는 없을 수 있으므로, optional
type responseBody struct {
	StatusCode    int         `json:"resultCode"`
	ErrorCode     int         `json:"errorCode,omitempty"`
	ResultMessage string      `json:"resultMessage"`
	ResultData    interface{} `json:"resultData,omitempty"`
}

// Controller에 대한 공통 Response 모델
func response(server echo.Context, status, errorcode int, message string, data ...interface{}) error {
	response := responseBody{
		StatusCode:    status,
		ErrorCode:     errorcode,
		ResultMessage: message,
		ResultData:    data[0],
	}
	return server.JSON(status, response)
}

func Healthcheck(server echo.Context) error {
	return response(server, http.StatusOK, 0, "success", map[string]string{
		"status": "running",
		"health": "service is health",
	}, map[string]string{
		"version": "1.0.0",
	})
}
