package main

import (
	"fmt"
	"github.com/roisin087/gorilla/mux"
	"github.com/roisin087/server/server"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	uc := server.NewUserController()
	fmt.Println("Attempttostartserveratport....8181")
	r.HandleFunc("/", uc.WelcomeHandler)
	//map handlers to URL
	r.HandleFunc("/Users", uc.GetUsersHandler).Methods("GET")
	r.HandleFunc("/Users/{id}", uc.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/Users/", uc.UserCreate).Methods("POST")
	http.Handle("/", r) //register routes with net/http

	server.Startserver()

}
