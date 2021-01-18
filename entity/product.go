package entity

type Product struct {
	Id       string `bson:"_id"`
	Name     string
	Price    int64
	Quantity int32
	Images   []ProductImage `bson:"images"`
}

type ProductImage struct {
	Name string
	Path string
}
