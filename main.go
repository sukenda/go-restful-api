package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sukenda/go-restful-api/config"
	"github.com/sukenda/go-restful-api/controller"
	"github.com/sukenda/go-restful-api/exception"
	"github.com/sukenda/go-restful-api/repository"
	"github.com/sukenda/go-restful-api/service"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	// Setup Repository
	productRepository := repository.NewProductRepository(database)
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	productService := service.NewProductService(&productRepository)
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	productController := controller.NewProductController(&productService)
	userController := controller.NewUserController(&userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup Routing
	productController.Route(app)
	userController.Route(app)

	// Start App
	err := app.Listen(":8080")
	exception.PanicIfNeeded(err)
}
