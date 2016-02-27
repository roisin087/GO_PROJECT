package main

import (
	"fmt"
	"github.com/roisin087/client/client/controller"
	"runtime"
	"strconv"
	"time"
)

func main() {

	ch := make(chan string)
	out := make(chan string)

	bch := make(chan string, 4)

	////synchronous////

	go producer(ch, 300*time.Millisecond, " producer 1")
	//go producer(ch, 100*time.Millisecond, " producer 2")
	//go producer(ch, 150*time.Millisecond, " producer 3")

	go consumer(out)

	fmt.Println(runtime.NumGoroutine())

	for i := range ch {
		time.Sleep(100 * time.Millisecond)
		out <- i

	}

	var input string
	fmt.Scanln(&input)

	/////////asynchronous

	go producer2(bch, 0*time.Millisecond, "post sent on buffered channel")

	for i := range bch {
		time.Sleep(100 * time.Millisecond)
		out <- i

	}

	fmt.Println(runtime.NumGoroutine())

}

func producer(ch chan string, d time.Duration, msg string) {

	for i := 1; i < 9; i++ {
		ch <- client.PostRequest(int32(i), "someone", "someone@live.ie") + "  received from " + msg
		fmt.Println("\n....post request sent from " + msg)
		time.Sleep(d)
	}

	close(ch)

}
func producer2(bch chan string, d time.Duration, msg string) {
	for i := 1; i < 9; i++ {
		bch <- client.GetResponse("http://localhost:8181/Users/"+strconv.Itoa(i)) + "  received from " + msg
		fmt.Println("\nget request sent from " + msg)
		time.Sleep(d)
	}
	close(bch)
}

func consumer(out chan string) {
	for res := range out {
		fmt.Println(res)
	}
	close(out)
}
