package server

type User struct {
	ID    int32  `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
type users struct {
	Users []User `json:"user"`
}
