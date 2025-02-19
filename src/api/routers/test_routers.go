package routers

import (
	"github.com/Masih-Ghasri/EchoStorm/api/handlers"
	"github.com/labstack/echo/v4"
)

func TestRouters(e *echo.Group) {
	h := handlers.NewTestHandler()

	e.GET("/test", h.Test)
	e.POST("/name", h.Query)
	e.POST("/bind", h.Bind)
}
