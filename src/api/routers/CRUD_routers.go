package routers

import (
	"github.com/Masih-Ghasri/EchoStorm/api/handlers"
	"github.com/labstack/echo/v4"
)

func CRUDRouters(e *echo.Group) {
	h := handlers.NewCRUDHandler()

	e.POST("/products", h.CreateProduct)
	e.GET("/products", h.GetProduct)
	e.GET("/products/:id", h.GetProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)
}
