package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

//expose REST API
func (uc UserController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	users := users{make([]user, 0)}

	for _, value := range usersMap {
		users.Users = append(users.Users, value)
		fmt.Println(value.Name)
		fmt.Println(len(users.Users))
	}

	data, error := json.Marshal(users)
	checkErrorAndPanic(error)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func (uc UserController) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome")
}
