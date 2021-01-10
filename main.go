package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/sukenda/go-restful-api/config"
	"github.com/sukenda/go-restful-api/controller"
	"github.com/sukenda/go-restful-api/exception"
	"github.com/sukenda/go-restful-api/repository"
	"github.com/sukenda/go-restful-api/service"
	"os"
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

	file, err := os.OpenFile(configuration.Get("LOG_DIR"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	exception.PanicIfNeeded(err)
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${body}\n",
		Output: file,
	}))

	// Setup Routing without Authorization
	userController.Route(app)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(configuration.Get("JWT_SECRET")),
	}))

	// Setup Routing with Authorization
	productController.Route(app)

	// Start App
	err = app.Listen(":8080")
	exception.PanicIfNeeded(err)
}
