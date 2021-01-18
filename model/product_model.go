package model

type CreateProductRequest struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Price    int64          `json:"price"`
	Quantity int32          `json:"quantity"`
	Images   []ProductImage `json:"images"`
}

type CreateProductResponse struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Price    int64          `json:"price"`
	Quantity int32          `json:"quantity"`
	Images   []ProductImage `json:"images"`
}

type GetProductResponse struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Price    int64          `json:"price"`
	Quantity int32          `json:"quantity"`
	Images   []ProductImage `json:"images"`
}

type ProductImage struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
