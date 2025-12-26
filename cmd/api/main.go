package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("Error initializing database")
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
}
