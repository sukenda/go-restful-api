package controller

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sukenda/go-restful-api/entity"
	"github.com/sukenda/go-restful-api/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

const JSON = "application/json"

func TestProductController_Save(t *testing.T) {
	productRepository.DeleteAll()

	createProductRequest := model.CreateProductRequest{
		Name:     "Test Product",
		Price:    10000,
		Quantity: 1000,
	}

	requestBody, _ := json.Marshal(createProductRequest)

	request := httptest.NewRequest("POST", "/api/products", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", JSON)
	request.Header.Set("Accept", JSON)
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI5N2RlNTQ3Ny1jY2RiLTQ2MWMtYjUyMS05ZmIwZTczNjBlOWUiLCJ1c2VybmFtZSI6InN1a2VuZGEiLCJyb2xlIjoiQWRtaW4iLCJleHAiOjE2MTA3ODUzMjd9.qwMSzCCC1SIMIOcvuIudzvFBoXcmcp63c7nIGlQSZLc")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createProductResponse := model.CreateProductResponse{}
	_ = json.Unmarshal(jsonData, &createProductResponse)
	assert.NotNil(t, createProductResponse.Id)
	assert.Equal(t, createProductRequest.Name, createProductResponse.Name)
	assert.Equal(t, createProductRequest.Price, createProductResponse.Price)
	assert.Equal(t, createProductRequest.Quantity, createProductResponse.Quantity)
}

func TestProductController_Update(t *testing.T) {

}

func TestProductController_Find(t *testing.T) {
	productRepository.DeleteAll()

	product := entity.Product{
		Id:       uuid.New().String(),
		Name:     "Sample Product",
		Price:    10000,
		Quantity: 1000,
	}

	productRepository.Insert(product)

	request := httptest.NewRequest("GET", "/api/products", nil)
	request.Header.Set("Accept", JSON)
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI5N2RlNTQ3Ny1jY2RiLTQ2MWMtYjUyMS05ZmIwZTczNjBlOWUiLCJ1c2VybmFtZSI6InN1a2VuZGEiLCJyb2xlIjoiQWRtaW4iLCJleHAiOjE2MTA3ODUzMjd9.qwMSzCCC1SIMIOcvuIudzvFBoXcmcp63c7nIGlQSZLc")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.([]interface{})
	containsProduct := false

	for _, data := range list {
		jsonData, _ := json.Marshal(data)
		getProductResponse := model.GetProductResponse{}
		_ = json.Unmarshal(jsonData, &getProductResponse)
		if getProductResponse.Id == product.Id {
			containsProduct = true
		}
	}

	assert.True(t, containsProduct)
}

func TestProductController_FindById(t *testing.T) {

}

func TestProductController_Delete(t *testing.T) {

}