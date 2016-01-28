package server

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type users struct {
	Users []user `json:"user"`
}
