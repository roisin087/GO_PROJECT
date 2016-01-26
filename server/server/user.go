package server

type user struct {
	Name string `xml:"name"`
}
type users struct {
	Users []user `xml:"user"`
}

var usersMap = make(map[string]user)
