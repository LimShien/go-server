package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func initDb() {
	connStr := "user=user password=password dbname=database sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
	createTable(db)
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id SERIAL PRIMARY KEY,
			item_name VARCHAR(255),
			item_notes TEXT,
			item_default_price NUMERIC,
			item_default_duration INT
		);

		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL PRIMARY KEY,
			customer_name VARCHAR(255) NOT NULL,
			customer_number VARCHAR(20) NOT NULL,
			customer_note TEXT,
			appointment_history JSONB
		);

		CREATE TABLE IF NOT EXISTS appointments (
			id SERIAL PRIMARY KEY,
			customer_id INTEGER REFERENCES customers(id),
			appointment_time TIMESTAMP,
			appointment_item VARCHAR(255),
			appointment_duration INT
		);
		
    `)
	return err
}
