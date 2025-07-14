package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
