package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID    int32
	Name  string
	Email string
}

func GetResponse(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Unexpected Status code", res.StatusCode)
		return res.Status

	}

	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
		return res.Status
	}
	res.Body.Close()
	return res.Status

}

func PostRequest(id int32, name string, email string) string {
	u := &User{id, name, email}
	buf, _ := json.Marshal(u)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://localhost:8181/Users/", "text/json", body)
	r.Body.Close()
	return "Post" + r.Status
}
