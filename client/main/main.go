package main

import (
	"client/client/controller"
	//"client/client/model"
	"fmt"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {
	// call publish method that will block the goroutine which sends a http request to the api
	go client.Publish("A goroutine starts a new thread of execution\n.", 2*time.Second)
	go client.Publish("Another goroutine starts should execute before other thread is finished\n", 5*time.Second)
	go fmt.Println("Thread should execute before I leave.\n")

	// Wait for the news to be published.
	time.Sleep(10 * time.Second)

	fmt.Println("\n\nTen seconds later: I’m leaving now.")
	fmt.Println("\nCalling api for user with id of 1")
	client.GetResponse("http://localhost:8181/Users/1")
	client.CheckConnection()

}
