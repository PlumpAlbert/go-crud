package core

import (
	"database/sql"
	"html/template"
	"log"
)

type AppConfig struct {
	Database  *sql.DB
	Templates *template.Template
}

var Config AppConfig

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1)/testdb?parseTime=true")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Check DB connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil
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
