package service

import "github.com/sukenda/go-restful-api/model"

type ProductService interface {
	Create(request model.CreateProductRequest) (response model.CreateProductResponse)

	List() (responses []model.GetProductResponse)

}