package server

import (
	"encoding/json"
	"fmt"
	"github.com/roisin087/gorilla/context"
	"github.com/roisin087/gorilla/mux"
	//"github.com/roisin087/gorilla/schema"
	"github.com/roisin087/server/database"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type key int

const mykey key = 0

type (
	// UserController represents the controller for operating on the User resource
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

//expose REST API
func (uc UserController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	sessionCopy := database.GetSession()
	defer sessionCopy.Close()

	users := sessionCopy.DB("MyDB").C("MyCollection")
	var results []User
	err := users.Find(bson.M{}).All(&results)
	if err != nil {
		panic(err)
	}
	for _, res := range results {
		fmt.Printf("User: %s|%s\n", res.Name, res.Email)
	}
	data, error := json.Marshal(results)
	checkErrorAndPanic(error)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func (uc UserController) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome")

}

func (uc UserController) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {

	//Vars used to extract id parameter
	vars := mux.Vars(r)
	id := vars["id"]
	SetMyKey(r, id)

	SearchByIdHandler(w, r)
}

func SearchByIdHandler(w http.ResponseWriter, r *http.Request) {

	sessionCopy := database.GetSession()
	defer sessionCopy.Close()
	i64, _ := strconv.ParseInt(GetMyKey(r), 10, 64)
	i := int32(i64)
	fmt.Println(i)
	fmt.Println("Gorilla context Key " + GetMyKey(r))

	user := User{}

	err := sessionCopy.DB("MyDB").C("MyCollection").Find(bson.M{"id": i}).One(&user)

	if err != nil {
		fmt.Println("no user with that id")
		data, error := json.Marshal("error no user can be found with an id of " + GetMyKey(r))
		checkErrorAndPanic(error)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {

		fmt.Println("User", user)
		data, error := json.Marshal(user)
		checkErrorAndPanic(error)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

}

//var decoder = schema.NewDecoder()

func (uc UserController) UserCreate(w http.ResponseWriter, r *http.Request) {
	sessionCopy := database.GetSession()
	defer sessionCopy.Close()

	//not working through schema
	/*
		var user User
		err := r.ParseForm()
		if err != nil {
			log.Println("Form could not be parsed")
		}
		err2 := decoder.Decode(&user, r.PostForm)
		if err2 != nil {
			log.Println("Got error decoding form: ", err2)
		}
		fmt.Println("User:" + user.Name)
	*/
	//end gorilla schema

	i := GetMyKey(r)
	fmt.Println("gorilla context not available here outside of request key empty" + i)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Println(string(body))
	var u User
	err = json.Unmarshal([]byte(string(body)), &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Println("Name" + u.Name)

	users := sessionCopy.DB("MyDB").C("MyCollection")
	err = users.Insert(User{ID: u.ID, Name: u.Name, Email: u.Email})
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("created")

}

func GetMyKey(r *http.Request) string {
	if rv := context.Get(r, mykey); rv != nil {
		return rv.(string)
	}
	return ""
}

func SetMyKey(r *http.Request, val string) {
	context.Set(r, mykey, val)
}
