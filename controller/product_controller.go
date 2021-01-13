package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sukenda/go-restful-api/exception"
	"github.com/sukenda/go-restful-api/model"
	"github.com/sukenda/go-restful-api/service"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/api/products", controller.Save)
	app.Put("/api/products", controller.Update)
	app.Get("/api/products", controller.Find)
	app.Get("/api/products/:id", controller.FindById)
	app.Delete("/api/products/:id", controller.Delete)
}

func (controller *ProductController) Save(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.Save(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) Update(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.ProductService.Update(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) Find(c *fiber.Ctx) error {
	responses := controller.ProductService.Find()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *ProductController) FindById(c *fiber.Ctx) error {
	var id = c.Params("id")
	responses := controller.ProductService.FindById(id)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *ProductController) Delete(c *fiber.Ctx) error {
	var id = c.Params("id")
	controller.ProductService.Delete(id)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Delete success",
	})
}
