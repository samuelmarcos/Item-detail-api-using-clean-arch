package config

import (
	"os"
	"testing"
)

func TestEnvironmentLoadsEnvVars(t *testing.T) {
	os.Setenv("MYSQL_USER", "testuser")
	os.Setenv("MYSQL_PASSWORD", "testpass")
	os.Setenv("MYSQL_HOST", "localhost")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_DATABASE", "testdb")

	cfg := Environment()

	if cfg.DB_USER != "testuser" {
		t.Errorf("Expected DB_USER to be 'testuser', got '%s'", cfg.DB_USER)
	}
	if cfg.DB_PASSWORD != "testpass" {
		t.Errorf("Expected DB_PASSWORD to be 'testpass', got '%s'", cfg.DB_PASSWORD)
	}
	if cfg.DB_HOST != "localhost" {
		t.Errorf("Expected DB_HOST to be 'localhost', got '%s'", cfg.DB_HOST)
	}
	if cfg.DB_PORT != "3306" {
		t.Errorf("Expected DB_PORT to be '3306', got '%s'", cfg.DB_PORT)
	}
	if cfg.DB_DATABASE != "testdb" {
		t.Errorf("Expected DB_DATABASE to be 'testdb', got '%s'", cfg.DB_DATABASE)
	}
}
