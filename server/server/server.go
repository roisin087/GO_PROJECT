package server

import (
	"fmt"
	"net/http"
)

//var usersMap = make(map[string]user)
func Startserver() {
	fmt.Println("Server running")
	u := user{Name: "Roisin"}
	usersMap["name"] = u
	err := http.ListenAndServe(":8181", nil)
	checkErrorAndPanic(err)
}

func checkErrorAndPanic(err error) {

	if err != nil {
		panic(err)
	}

}
