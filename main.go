package main

import (
	"authenncrm/config"
	"authenncrm/internal/handlers"
	"authenncrm/internal/infrastructures"
	"authenncrm/internal/repositories"
	"authenncrm/internal/services"
	"fmt"
	"log"

	"github.com/fxamacker/cbor/v2"
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

	authService := services.NewAuthenticationService(userRepository, &cfg)
	authHandler := handlers.NewAuthenticationHandler(authService)

	app := fiber.New(fiber.Config{
		CBOREncoder: cbor.Marshal,
		CBORDecoder: cbor.Unmarshal,
	})

	app.Get("/health-check", healthCheckHandler.CheckHealth)

	apiV1 := app.Group("/api/v1")

	authenticationGroup := apiV1.Group("/auth")
	authenticationGroup.Post("/login", authHandler.Login)
	// authenticationGroup.Post("/refresh-token", "")
	// authenticationGroup.Post("/logout", "", authentication.AuthMiddleware)
	// authenticationGroup.Post("/register", "")
	// authenticationGroup.Post("/reset-password", "")

	userGroup := apiV1.Group("/users")
	userGroup.Post("", userHandler.CreateUser)
	userGroup.Get("/:userId", userHandler.GetUserById)
	userGroup.Get("", userHandler.GetAllUsers)
	userGroup.Put("/:userId", userHandler.UpdateUser)
	userGroup.Delete("/:userId", userHandler.DeleteUser)

	log.Printf("Now server is running. Checking %v:%v/health-check", cfg.ServerHost, cfg.ServerPort)
	app.Listen(fmt.Sprintf(":%v", cfg.ServerPort))
}
