package main

import (
	"bytes"
	"client/client/controller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"time"
)

func main() {

	client.CheckConnection()

	type User struct {
		ID    int32
		Name  string
		Email string
	}

	u := &User{1, "Adam", "adam@gmail.com"}
	buf, _ := json.Marshal(u)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://localhost:8181/User/", "text/json", body)
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))

	fmt.Println("\nCalling api for user with id of 1")
	client.GetResponse("http://localhost:8181/User/1")

	// call publish method that will block the goroutine which sends a http request to the api
	/*
		go client.Publish("A goroutine starts a new thread of execution\n.", 2*time.Second)
		go client.Publish("Another goroutine starts should execute before other thread is finished\n", 5*time.Second)
		go fmt.Println("Thread should execute before I leave.\n")

		// Wait for the news to be published.
		time.Sleep(10 * time.Second)

		fmt.Println("\n\nTen seconds later: I’m leaving now.")
	*/
	/*
		url := "http://localhost:8181/User/"
		fmt.Println("URL:>", url)

		var jsonStr = []byte(`{"id":5,"name":"Joe","email":"h@gmail.com"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))*/

}
