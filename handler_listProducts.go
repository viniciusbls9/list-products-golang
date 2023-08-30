package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handlerListProducts(c echo.Context) error {
	var result []Product // Change to the appropriate type to match your data structure
	db, err := handlerOpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare and execute the query
	rows, err := db.Query("SELECT Name, Price FROM products")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Name, &product.Price); err != nil {
			return err
		}
		result = append(result, product)
	}

	return c.JSON(http.StatusOK, result)
}
