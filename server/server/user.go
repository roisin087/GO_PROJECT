package server

type user struct {
	//ID    bson.ObjectId `bson:"_id" json:"id"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type users struct {
	Users []user `json:"user"`
}
