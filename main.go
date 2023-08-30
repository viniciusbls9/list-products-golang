package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	product := Product{Name: "Iphone", Price: 1000}
	persistProduct(product)
	w.Write([]byte("Hello World"))
}

func persistProduct(product Product) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	stmt, err := db.Prepare("Insert into products(name, price) values($1, $2)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(product.Name, product.Price)
}
