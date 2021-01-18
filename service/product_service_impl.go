package service

import (
	"github.com/sukenda/go-restful-api/entity"
	"github.com/sukenda/go-restful-api/model"
	"github.com/sukenda/go-restful-api/repository"
	"github.com/sukenda/go-restful-api/validation"
)

func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		repository: *productRepository,
	}
}

type productServiceImpl struct {
	repository repository.ProductRepository
}

func (service *productServiceImpl) Save(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.ValidateProduct(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.repository.Insert(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (service *productServiceImpl) Update(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.ValidateProduct(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.repository.Update(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (service *productServiceImpl) Delete(id string) {
	service.repository.Delete(id)
}

func (service *productServiceImpl) Find() (responses []model.GetProductResponse) {
	products := service.repository.FindAll()

	for _, product := range products {
		var images []model.ProductImage
		for _, image := range product.Images {
			images = append(images, model.ProductImage{
				Name: image.Name,
				Path: image.Path,
			})
		}

		responses = append(responses, model.GetProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return responses
}

func (service *productServiceImpl) FindById(id string) (response model.CreateProductResponse) {
	product := service.repository.FindById(id)
	var images []model.ProductImage

	for _, image := range product.Images {
		images = append(images, model.ProductImage{
			Name: image.Name,
			Path: image.Path,
		})
	}

	return model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Images:   images,
	}
}

func (service *productServiceImpl) Upload(id string, images []model.ProductImage) (response model.CreateProductResponse) {
	var params []entity.ProductImage
	for _, image := range images {
		params = append(params, entity.ProductImage{
			Name: image.Name,
			Path: image.Path,
		})
	}

	product := service.repository.Upload(id, params)
	return model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Images:   images,
	}
}
