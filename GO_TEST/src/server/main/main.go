package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"server/server"
)

func main() {
	r := mux.NewRouter()
	uc := server.NewUserController()
	fmt.Println("Attempttostartserveratport....8181")
	fmt.Println("map contents")
	r.HandleFunc("/", uc.WelcomeHandler)
	//map handlers to URL
	r.HandleFunc("/Users", uc.GetUsersHandler).Methods("GET")
	r.HandleFunc("/Users/{id}", uc.GetUserByIDHandler).Methods("GET")
	http.Handle("/", r) //register routes with net/http
	server.Startserver()

}
