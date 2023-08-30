package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handlerCreateProduct(c echo.Context) error {
	product := Product{}
	c.Bind(&product)

	err := handlerPersistProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, product)
}
