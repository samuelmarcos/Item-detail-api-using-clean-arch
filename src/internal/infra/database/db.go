package database

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite", "product.db")
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(
		context.Background(),
		`DROP TABLE IF EXISTS product_details;
		 CREATE TABLE product_details (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			brand TEXT,
			model TEXT,
			color TEXT,
			category TEXT,
			images TEXT, -- JSON array of strings
			price REAL,
			original_price REAL,
			discount_percent REAL,
			stock INTEGER,
			seller_info TEXT, -- JSON object
			warranty TEXT,
			payment_options TEXT, -- JSON array of strings
			specs TEXT, -- JSON object
			related_products TEXT, -- JSON array of strings
			rating REAL,
			review_count INTEGER,
			free_shipping BOOLEAN,
			purchase_options TEXT, -- JSON array of strings
			highlights TEXT, -- JSON array of strings
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
