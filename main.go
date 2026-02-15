package main

import (
	"authenncrm/config"
	"authenncrm/internal/infrastructures"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	cfg := config.ReadInConfig()

	_, err := infrastructures.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	app := fiber.New()

	// db.AutoMigrate(
	// 	&entities.User{},
	// 	&entities.Role{},
	// 	&entities.PointLogs{},
	// )

	log.Println("Database migration completed successfully")

	//run server

	app.Listen(fmt.Sprintf(":%v", cfg.AppPort))

}
