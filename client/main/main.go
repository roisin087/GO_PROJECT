package main

import (
	"client/client/controller"

	"fmt"
	//"io/ioutil"
	//"math/rand"
	//"strconv"
	"time"
)

func main() {

	client.CheckConnection()

	//fmt.Println("\nCalling api for user with id of 1")
	//client.GetResponse("http://localhost:8181/Users/1")

	// call publish method that will block the goroutine which sends a http request to the api
	/*
		go client.Publish("A goroutine starts a new thread of execution\n.", 2*time.Second)
		go client.Publish("Another goroutine starts should execute before other thread is finished\n", 5*time.Second)
		go fmt.Println("Thread should execute before I leave.\n")

		// Wait for the news to be published.
		time.Sleep(10 * time.Second)

		fmt.Println("\n\nTen seconds later: I’m leaving now.")
	*/

	var channel1 chan string = make(chan string)
	var channel2 chan string = make(chan string)
	var channel3 chan string = make(chan string)
	for i := 0; i < 2; i++ {
		go sendPostRequest(channel1, "channel 1", 100*time.Millisecond)
	}
	for i := 0; i < 2; i++ {
		go sendGetRequest(channel2, "channel 2", 1000*time.Millisecond)
	}
	for i := 0; i < 2; i++ {
		go sendPostRequest(channel3, "channel 3", 500*time.Millisecond)
	}
	for i := 0; i < 2; i++ {
		go sendGetRequest(channel2, "channel 2", 800*time.Millisecond)
	}
	go receiveFromChannel(channel1, "Receiving from channel 1: ", 0*time.Millisecond)
	go receiveFromChannel(channel2, "Receiving from channel 2: ", 0*time.Millisecond)
	go receiveFromChannel(channel3, "Receiving from channel 3: ", 0*time.Millisecond)
	go receiveFromChannel(channel2, "Receiving from channel 3: ", 0*time.Millisecond)
	var input string
	fmt.Scanln(&input)

}

func sendPostRequest(channel chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	channel <- client.PostRequest(4, "nicola", "n@live.ie")
	fmt.Println("Post Request sent through " + msg)

}

func sendGetRequest(channel chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	channel <- client.GetResponse("http://localhost:8181/Users/1")
	fmt.Println("Get Request sent through " + msg)
}

func receiveFromChannel(channel <-chan string, msg string, delay time.Duration) { // returns receive-only channel of strings.
	time.Sleep(delay)
	for m := range channel {
		fmt.Println(msg + m)
	}

}
