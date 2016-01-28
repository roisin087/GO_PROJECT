package client

import (
	"fmt"
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
