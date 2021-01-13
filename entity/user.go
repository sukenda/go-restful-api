package entity

type User struct {
	Id       string `bson:"_id"`
	Username string
	Password string
	Email    string
	Phone    string
}
