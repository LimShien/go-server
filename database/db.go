package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			item_name VARCHAR(255),
			item_notes TEXT,
			item_default_price NUMERIC,
			item_default_duration INT
		);

		CREATE TABLE IF NOT EXISTS customers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customer_name VARCHAR(255) NOT NULL,
			customer_number VARCHAR(20) NOT NULL,
			customer_note TEXT,
			appointment_history JSONB
		);

		CREATE TABLE IF NOT EXISTS appointments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customer_id INTEGER REFERENCES customers(id),
			appointment_time TIMESTAMP,
			appointment_item VARCHAR(255),
			appointment_duration INT
		);
		
    `)
	return err
}

func CloseDB(db *sql.DB) {
	db.Close()
	fmt.Println("Database connection closed")
}
