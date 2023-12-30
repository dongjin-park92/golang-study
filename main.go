package main

import (
	"fmt"
	"golang-go/configs"
	"golang-go/loggers"
	"golang-go/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const (
	banner = `
	██████   ██████   ██████   ██████  
	██       ██    ██ ██       ██    ██ 
	██   ███ ██    ██ ██   ███ ██    ██ 
	██    ██ ██    ██ ██    ██ ██    ██ 
	 ██████   ██████   ██████   ██████  
	                                    
	`
)

func main() {
	conf := configs.GetConfig()
	startServer(conf)
}

func startServer(conf *configs.ViperConfig) {
	// 기본적인 Echo Framwork Instance 생성
	server := echo.New()
	server.Logger.SetLevel(log.INFO)
	loggers.InitLogger(server)

	// Echo Framwork에 대한 Middleware 설정
	// Logger, Recover(PANIC 복구)를 기본으로 사용
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// 전체 Origin Domain에 대해서 허용, Production에서는 FE 도메인으로 조정하는것을 권고
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET},
	}))

	// Echo Framwork의 기본 Banner를 삭제하고, Custom Banner 출력
	server.HideBanner = true
	fmt.Println(banner)

	// WebServer Start
	routers.InitRoute(server)
	server.Logger.Fatal(server.Start(conf.GetString("server.port")))

}
