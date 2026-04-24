package main

import (
	"log"

	"github.com/harshit14100/go-todo/config"
	"github.com/harshit14100/go-todo/database"
	"github.com/harshit14100/go-todo/server"
)

func main() {
	config.LoadConfig()

	err := database.ConnectandMigrate(
		config.GetEnv("DB_HOST", "-"),
		config.GetEnv("DB_PORT", "-"),
		config.GetEnv("DB_NAME", "-"),
		config.GetEnv("DB_USER", "-"),
		config.GetEnv("DB_PASSWORD", "password"),
		database.SSLMode(config.GetEnv("DB_SSLMode", "disable")),
	)
	if err != nil {
		log.Fatal(err)
	}

	r := server.NewServer()
	port := config.GetEnv("PORT", "8080")
	r.Run(":" + port)

}
