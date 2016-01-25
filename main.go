	package main

	import (
		"encoding/json"
		"encoding/xml"
		"fmt"
		"net/http"
		"strings"
	)

	func checkErrorAndPanic(err error) {

		if err != nil {
			panic(err)
		}

	}

	type user struct {
		Name string `xml:"name"`
	}
	type users struct {
		Users []user `xml:"user"`
	}

	var usersMap map[string]user

	func main() {
		usersMap = make(map[string]user)
		u := user{Name: "Roisin"}
		usersMap["name"] = u
		fmt.Println("Attempttostartserveratport....8181")
		http.HandleFunc("/", welcomeHandler)
		//map handlers to URL
		http.HandleFunc("/getRegUser", getRegisteredUsersHandler)
		http.HandleFunc("/rest/json/getRegUser", getRegisteredUsersHandlerRestAPI)
		http.HandleFunc("/rest/xml/getRegUser", getRegisteredUsersHandlerRestAPI)
		err := http.ListenAndServe(":8181", nil)
		checkErrorAndPanic(err)
	}

	func welcomeHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Server running")
	}


	//define handler
	func​ getRegisteredUsersHandler(w http.ResponseWriter, r *http.Request) {

		regUser := ""

		for​ _, user := range​ usersMap {

			if​(len(user.Name)>0){

			regUser = regUser + "\n" + user.Name

			}

		}

		w.Write([]byte("Registered Users are:" + regUser))

	}
	//expose REST API
	func getRegisteredUsersHandlerRestAPI(w http.ResponseWriter, r *http.Request) {

		users := users{make([]user, 0)}

		for _, value := range usersMap {
			users.Users = append(users.Users, value)
			fmt.Println(value.Name)
			fmt.Println(len(users.Users))

		}
		if strings.Contains(r.URL.Path, "/json/") {
			data, error := json.Marshal(users)
			checkErrorAndPanic(error)
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		} else if strings.Contains(r.URL.Path, "/xml/") {
			data, error := xml.Marshal(users)
			checkErrorAndPanic(error)
			w.Header().Set("Content-Type", "application/xml")
			w.Write(data)
		}
	}
