package core

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"os"
)

type AppConfig struct {
	Database  *sql.DB
	Templates *template.Template
}

var Config AppConfig

func initDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbSsl := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbDatabase,
	)

	if dbSsl != "" {
		dsn += " sslmode=" + dbSsl
	}

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	// Check DB connection
	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func NewConfig() *AppConfig {
	Config = AppConfig{
		Database:  initDB(),
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	return &Config
}
