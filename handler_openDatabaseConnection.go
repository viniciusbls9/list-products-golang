package main

import (
	"database/sql"
)

func handlerOpenDatabaseConnection() (*sql.DB, error) {
	dbURL := handlerGetEnv()

	// Open a database connection
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
