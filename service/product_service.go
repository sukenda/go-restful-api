package service

import "github.com/sukenda/go-restful-api/model"

type ProductService interface {
	Save(request model.CreateProductRequest) (response model.CreateProductResponse)

	Find() (responses []model.GetProductResponse)
}
