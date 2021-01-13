# Go Restful Api

Spring restful api build for learning go using fiber and mongodb

# API Spec

## Register account
Request :

- Method : POST
- Endpoint : `/api/signup`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
  "username": "string",
  "password": "string",
  "email": "string",
  "phone": "string"
}
```

Response :

```json 
{
  "code": "int",
  "status": "string",
  "data": {
    "id": "string",
    "username": "string",
    "password": "string",
    "email": "string",
    "phone": "string"
  }
}
```

## Login

Request :

- Method : POST
- Endpoint : `/api/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
  "username": "string",
  "password": "string"
}
```

Response :

```json 
{
  "code": "int",
  "status": "string",
  "data": {
    "access_token": "string",
    "user": {
      "id": "string",
      "username": "string",
      "email": "string",
      "phone": "string"
    }
  }
}
```

## Create Product

Request :

- Method : POST
- Endpoint : `/api/products`
- Header :
    - Authorization: Bearer xxxxx
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
  "name": "string",
  "price": "int64",
  "quantity": "int32"
}
```

Response :

```json
{
  "code": "int",
  "status": "string",
  "data": {
    "id": "string",
    "name": "string",
    "price": "int64",
    "quantity": "int32"
  }
}
```

## Get Product

Request :

- Method : GET
- Endpoint : `/api/products`
- Header :
    - Authorization: Bearer xxxxx
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
  "code": "int",
  "status": "string",
  "data": [
    {
      "id": "string",
      "name": "string",
      "price": "int64",
      "quantity": "int32"
    },
    {
      "id": "string",
      "name": "string",
      "price": "int64",
      "quantity": "int32"
    }
  ]
}
```

# Reference
https://github.com/khannedy/golang-clean-architecture