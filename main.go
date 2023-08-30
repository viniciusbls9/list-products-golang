package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	e := echo.New()
	e.GET("/product", handlerListProducts)
	e.POST("/product", handlerCreateProduct)
	e.Logger.Fatal(e.Start(":8080"))
}
