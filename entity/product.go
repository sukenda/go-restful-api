package entity

type Product struct {
	Id       string `bson:"_id"`
	Name     string
	Price    int64
	Quantity int32
}
