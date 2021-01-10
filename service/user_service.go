package service

import (
	"github.com/sukenda/go-restful-api/model"
)

type UserService interface {
	Insert(request model.CreateUserRequest) (response model.CreateUserResponse)

	FindByUsername(username string) (response model.CreateUserResponse)

	Login(request model.CreateUserRequest) (response model.GetLoginResponse)

}
