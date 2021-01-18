package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sukenda/go-restful-api/exception"
	"github.com/sukenda/go-restful-api/model"
	"github.com/sukenda/go-restful-api/service"
)

const PRODUCTS = "/api/products"

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post(PRODUCTS, controller.Save)
	app.Put(PRODUCTS, controller.Update)
	app.Get(PRODUCTS, controller.Find)
	app.Get(PRODUCTS+"/:id", controller.FindById)
	app.Delete(PRODUCTS+"/:id", controller.Delete)
	app.Post(PRODUCTS+"/:id/images", controller.Upload)
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
	// This sample for get token and parse
	/*bytes := c.Request().Header.Peek("Authorization")
	token := string(bytes)[7:]
	user, err := validation.ParseToken(token)
	exception.PanicIfNeeded(err)
	fmt.Println("ID ", user.Id)
	fmt.Println("Username ", user.Username)*/

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

func (controller *ProductController) Upload(c *fiber.Ctx) error {
	var id = c.Params("id")
	var images []model.ProductImage

	image1, err := c.FormFile("image1")
	exception.PanicIfNeeded(err)

	err1 := c.SaveFile(image1, fmt.Sprintf("./images/%s", image1.Filename))
	if err1 == nil {
		images = append(images, model.ProductImage{
			Name: image1.Filename,
			Path: fmt.Sprintf("./images/%s", image1.Filename),
		})
	}

	image2, err := c.FormFile("image2")
	exception.PanicIfNeeded(err)

	err2 := c.SaveFile(image2, fmt.Sprintf("./images/%s", image2.Filename))
	if err2 == nil {
		images = append(images, model.ProductImage{
			Name: image2.Filename,
			Path: fmt.Sprintf("./images/%s", image2.Filename),
		})
	}

	response := controller.ProductService.Upload(id, images)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
