package main

import (
	"fmt"
	"github.com/client/client/controller"
	"time"
)

func main() {

	client.CheckConnection()

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

func receiveFromChannel(channel <-chan string, msg string, delay time.Duration) { //prints from channel
	time.Sleep(delay)
	for m := range channel {
		fmt.Println(msg + m)
	}

}
