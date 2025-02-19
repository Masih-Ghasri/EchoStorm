package api

import (
	"github.com/Masih-Ghasri/EchoStorm/api/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	v1 := api.Group("/v1")
	{
		testRouter := v1.Group("/test")
		routers.TestRouters(testRouter)

		crudRouter := v1.Group("/crud")
		routers.CRUDRouters(crudRouter)
	}

	return e
}

func StartServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":1323"))
}
