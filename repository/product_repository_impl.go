package repository

import (
	"github.com/sukenda/go-restful-api/config"
	"github.com/sukenda/go-restful-api/entity"
	"github.com/sukenda/go-restful-api/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductRepository(database *mongo.Database) ProductRepository {
	return &productRepositoryImpl{
		Collection: database.Collection("products"),
	}
}

type productRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository productRepositoryImpl) Upload(id string, images []entity.ProductImage) (product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"images": images,
		},
	})
	exception.PanicIfNeeded(err)

	return repository.FindById(id)
}

func (repository productRepositoryImpl) Insert(product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":      product.Id,
		"name":     product.Name,
		"price":    product.Price,
		"quantity": product.Quantity,
	})
	exception.PanicIfNeeded(err)
}

func (repository productRepositoryImpl) Update(product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.UpdateOne(ctx, bson.M{"_id": product.Id}, bson.M{
		"$set": bson.M{
			"name":     product.Name,
			"price":    product.Price,
			"quantity": product.Quantity,
		},
	})
	exception.PanicIfNeeded(err)
}

func (repository productRepositoryImpl) FindAll() (products []entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		/*images := document["images"].([]primitive.A)
		var result []entity.ProductImage
		for _, image := range images {
			//imageMap := image.(bson.M)
			result = append(result, entity.ProductImage{
				Name: image["name"].(string),
				Path: image["path"].(string),
			})
		}*/

		products = append(products, entity.Product{
			Id:       document["_id"].(string),
			Name:     document["name"].(string),
			Price:    document["price"].(int64),
			Quantity: document["quantity"].(int32),
		})
	}

	return products
}

func (repository productRepositoryImpl) FindById(id string) (product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	err := repository.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	exception.PanicIfNeeded(err)

	return product
}

func (repository productRepositoryImpl) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}

func (repository productRepositoryImpl) Delete(id string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{"_id": id})
	exception.PanicIfNeeded(err)
}
