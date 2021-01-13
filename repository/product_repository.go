package repository

import "github.com/sukenda/go-restful-api/entity"

type ProductRepository interface {

	Insert(product entity.Product)

	Update(product entity.Product)

	FindById(id string) (product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()

	Delete(id string)
}
