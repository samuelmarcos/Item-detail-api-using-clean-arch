package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func NewDB(config DBConfig) (*DB, error) {
	user := config.User
	password := config.Password
	host := config.Host
	if host == "" {
		host = "localhost"
	}
	port := config.Port
	if port == "" {
		port = "3306"
	}
	dbName := config.Database

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
