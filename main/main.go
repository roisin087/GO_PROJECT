package main

import (
	"fmt"
	"github.com/roisin087/client/client/controller"
	"sync"
	"time"
)

func main() {

	ch := make(chan string)
	out := make(chan string)

	bch := make(chan string, 4)

	go producer(ch, 300*time.Millisecond, "post sent from producer 1")
	//	go producer(ch, 100*time.Millisecond, "post sent from producer 2")
	//	go producer(ch, 150*time.Millisecond, "post sent from producer 3")
	go consumer(out)

	for i := range ch {
		time.Sleep(100 * time.Millisecond)
		out <- i

	}

	var input string
	fmt.Scanln(&input)
	go producer2(bch, 0*time.Millisecond, "post sent on buffered producer 1")
	for i := range bch {
		time.Sleep(100 * time.Millisecond)
		out <- i
	}

}

func producer(ch chan string, d time.Duration, msg string) {

	for i := 0; i < 8; i++ {
		ch <- client.PostRequest(4, "nicola", "n@live.ie")
		fmt.Println(msg)
		time.Sleep(d)
	}
	close(ch)
}
func producer2(bch chan string, d time.Duration, msg string) {
	for i := 0; i < 8; i++ {
		bch <- client.GetResponse("http://localhost:8181/Users/1")
		fmt.Println(msg)
		time.Sleep(d)
	}
	close(bch)
}

func consumer(out chan string) {
	for res := range out {
		fmt.Println(res)
	}
}
