package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func validateUser(user User) error {
	if user.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is empty")
	}
	if user.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "email is empty")
	}
	if user.Age <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "age is less than zero")
	}
	if user.Country == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "country is empty")
	}
	return nil
}

func (h *TestHandler) Test(c echo.Context) error {
	err := c.JSON(http.StatusOK, echo.Map{
		"hello": "world",
		"stats": "working",
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *TestHandler) Query(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is empty")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "ok",
		"name":   name,
	})
}

func (h *TestHandler) Bind(c echo.Context) error {
	U := User{}
	err := c.Bind(&U)
	if err != nil {
		return err
	}

	if err := validateUser(U); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "ok",
		"data":   U,
	})

}
