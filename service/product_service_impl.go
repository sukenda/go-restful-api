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

func (service *productServiceImpl) Find() (responses []model.GetProductResponse) {
	products := service.repository.FindAll()
	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return responses
}
