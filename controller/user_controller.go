package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sukenda/go-restful-api/exception"
	"github.com/sukenda/go-restful-api/model"
	"github.com/sukenda/go-restful-api/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/api/signup", controller.Register)
	app.Post("/api/login", controller.Login)
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.UserService.Register(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.UserService.Login(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
