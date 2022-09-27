package app

import (
	"flybitch/app/controller"
	"flybitch/app/middleware"
)

func addRoutes() {
	api := e.Group("api")
	api.Use(middleware.Login)
	api.GET("/ping", controller.Ping)
}
