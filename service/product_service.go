package service

import "github.com/sukenda/go-restful-api/model"

type ProductService interface {
	Save(request model.CreateProductRequest) (response model.CreateProductResponse)

	Update(request model.CreateProductRequest) (response model.CreateProductResponse)

	Delete(id string)

	Find() (responses []model.GetProductResponse)

	FindById(id string) (response model.CreateProductResponse)

	Upload(id string, images []model.ProductImage) (response model.CreateProductResponse)
}
