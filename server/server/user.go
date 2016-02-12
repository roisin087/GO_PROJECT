package server

type User struct {
	ID    int32  `schema:"id" json:"id" bson:"id"`
	Name  string `schema: "name" json:"name" bson:"name"`
	Email string `schema: "email" json:"email" bson:"email"`
}
type users struct {
	Users []User `json:"user"`
}
