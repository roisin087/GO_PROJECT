package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/client/client/model"
	//	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID    int32
	Name  string
	Email string
}

func GetResponse(url string) string {
	// request http api
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Unexpected Status code", res.StatusCode)
		return res.Status

	}

	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
		return res.Status
	}

	//fmt.Printf("Body: %s\n", body)
	return res.Status

}

func CheckConnection() {
	if model.ConnectToMongo() {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connected")
	}
}

func PostRequest(id int32, name string, email string) string {
	u := &User{id, name, email}
	buf, _ := json.Marshal(u)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://localhost:8181/Users/", "text/json", body)
	//response, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(response))
	return "Post" + r.Status
}
