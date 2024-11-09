package main

import (
	"ecom/cmd/api"
	"ecom/db"
	"ecom/internal/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.Envs.User,
		config.Envs.Password,
		config.Envs.Name,
		config.Envs.Host,
		config.Envs.Port,
		config.Envs.SSLMode,
	)

	db, err := db.NewPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

	checkDatabase(db)

	server := api.NewAPIServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func checkDatabase(db *sqlx.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("database connected")
}
