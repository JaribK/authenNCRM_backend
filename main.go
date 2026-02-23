package main

import (
	"authenncrm/config"
	"authenncrm/internal/handlers"
	"authenncrm/internal/infrastructures"
	"authenncrm/internal/repositories"
	"authenncrm/internal/services"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	cfg := config.ReadInConfig()

	db, err := infrastructures.ConnectDatabase(&cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	healthCheckHandler := handlers.NewHealthCheckHandler()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New()

	app.Get("/health-check", healthCheckHandler.CheckHealth)

	userGroup := app.Group("/users")
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Get("/:id", userHandler.GetUserById)
	userGroup.Get("/", userHandler.GetAllUsers)
	userGroup.Put("/:id", userHandler.UpdateUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)

	log.Printf("Now server is running. Checking %v:%v/health-check", cfg.ServerHost, cfg.ServerPort)
	app.Listen(fmt.Sprintf(":%v", cfg.ServerPort))
}
