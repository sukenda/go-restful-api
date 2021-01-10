package repository

import (
	"errors"
	"github.com/sukenda/go-restful-api/config"
	"github.com/sukenda/go-restful-api/entity"
	"github.com/sukenda/go-restful-api/exception"
	"github.com/sukenda/go-restful-api/validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepositoryImpl{Collection: database.Collection("users")}
}

type userRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository userRepositoryImpl) Insert(user entity.User) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.FindByUsername(user.Username)
	if err == nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "User exist",
		})
	}

	_, err = repository.Collection.InsertOne(ctx, bson.M{
		"_id":      user.Id,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"phone":    user.Phone,
	})
	exception.PanicIfNeeded(err)

	return nil
}

func (repository userRepositoryImpl) FindByUsername(username string) (user entity.User, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{"username": username})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	// Return 1 or 0
	if len(documents) == 0 {
		return user, errors.New("USER_NOT_FOUND")
	}

	return entity.User{
		Id:       documents[0]["_id"].(string),
		Username: documents[0]["username"].(string),
		Password: documents[0]["password"].(string),
		Email:    documents[0]["email"].(string),
		Phone:    documents[0]["phone"].(string),
	}, nil
}

func (repository userRepositoryImpl) Login(username, password string) (user entity.User, err error) {
	_, cancel := config.NewMongoContext()
	defer cancel()

	user, err = repository.FindByUsername(username)
	exception.PanicIfNeeded(err)

	match, _ := validation.ValidatePassword(password, user.Password)
	if !match {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "Password not valid",
		})
	}

	return user, nil
}
