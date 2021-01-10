package repository

import "github.com/sukenda/go-restful-api/entity"

type ProductRepository interface {
	Insert(product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()
}
