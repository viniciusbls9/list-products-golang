package main

func handlerPersistProduct(product Product) error {
	db, err := handlerOpenDatabaseConnection()
	stmt, err := db.Prepare("Insert into products(name, price) values($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}
