package client

import (
	"client/client/model"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("ALERT:", text)
		GetResponse("http://localhost:8181/Users")

	}() // Note the parentheses. We must call the anonymous function.
}

func GetResponse(url string) {
	// request http api
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// read body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	fmt.Printf("Body: %s\n", body)

}

type user struct {
	ID    string
	Name  string
	Email string
}

//not enough arguments to return error ??
/*
// GetBSON implements bson.Getter.
func (u User) GetBSON() (interface{}, error) {

	return struct {
		ID    string `json:"id" bson:"id"`
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
	}{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
*/
// SetBSON implements bson.Setter.
func (u *user) SetBSON(raw bson.Raw) error {

	decoded := new(struct {
		ID    string `json:"id" bson:"id"`
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
	})

	bsonErr := raw.Unmarshal(decoded)

	if bsonErr == nil {
		u.ID = decoded.ID
		u.Name = decoded.Name
		u.Email = decoded.Email
		return nil
	} else {
		return bsonErr
	}
}

func CheckConnection() {
	if model.ConnectToMongo() {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connected")
	}
}
