package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/sukenda/go-restful-api/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestUserController_Register(t *testing.T) {
	createUserRequest := model.CreateUserRequest{
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
		Phone:    "085624556889",
	}

	requestBody, _ := json.Marshal(createUserRequest)
	request := httptest.NewRequest("POST", "/api/signup", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", JSON)
	request.Header.Set("Accept", JSON)

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createUserResponse := model.CreateUserResponse{}
	_ = json.Unmarshal(jsonData, &createUserResponse)

	assert.NotNil(t, createUserResponse.Id)

}

func TestUserController_Login(t *testing.T) {
	createUserRequest := model.CreateUserRequest{
		Username: "sukenda",
		Password: "sukenda",
	}

	requestBody, _ := json.Marshal(createUserRequest)

	request := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", JSON)
	request.Header.Set("Accept", JSON)

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	loginResponse := model.GetLoginResponse{}
	_ = json.Unmarshal(jsonData, &loginResponse)

	assert.NotNil(t, loginResponse.AccessToken)
	assert.NotNil(t, loginResponse.User)

}
