package main

import (
	"client/client"
	"fmt"
	//	"log"
	"time"
)

func main() {

	client.Publish("A goroutine starts a new thread of execution.", 5*time.Second)
	fmt.Println("Let’s hope the news will published before I leave.")

	// Wait for the news to be published.
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I’m leaving now.")
	fmt.Println("Calling api for user with id of 1")
	client.GetResponse("http://localhost:8181/Users?id=1")

}
