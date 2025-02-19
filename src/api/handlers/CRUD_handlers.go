package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CRUDHandler struct{}

func NewCRUDHandler() *CRUDHandler {
	return &CRUDHandler{}
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product = []Product{
	{ID: 1, Name: "Apple", Price: 20},
	{ID: 2, Name: "Mango", Price: 100},
	{ID: 3, Name: "Pineapple", Price: 150},
}
var currentID = 4

func (h *CRUDHandler) CreateProduct(c echo.Context) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	product.ID = currentID
	currentID++
	products = append(products, *product)
	return c.JSON(http.StatusCreated, product)
}
func (h *CRUDHandler) GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func (h *CRUDHandler) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, product := range products {
		if product.ID == id {
			return c.JSON(http.StatusOK, product)
		}
	}
	return c.String(http.StatusNotFound, "Product not found")
}

func (h *CRUDHandler) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	for i, p := range products {
		if p.ID == id {
			product.ID = id
			products[i] = *product
			return c.JSON(http.StatusOK, product)
		}
	}
	return c.String(http.StatusNotFound, "Product not found")
}

func (h *CRUDHandler) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.String(http.StatusNotFound, "Product not found")
}
