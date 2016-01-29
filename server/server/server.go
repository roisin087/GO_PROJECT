package server

import (
	"fmt"
	"net/http"
)

var usersMap = make(map[string]user)

func Startserver() {
	fmt.Println("Server running")

	u := user{ID: "1", Name: "Roisin", Email: "roisin@gmail.com"}
	u2 := user{ID: "2", Name: "Michelle", Email: "h@gmail.com"}
	usersMap[u.ID] = u
	usersMap[u2.ID] = u2
	err := http.ListenAndServe(":8181", nil)
	checkErrorAndPanic(err)
}

func checkErrorAndPanic(err error) {

	if err != nil {
		panic(err)
	}

}
