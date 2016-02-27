package main

import (
	"fmt"
	"github.com/roisin087/client/client/controller"
	"strconv"
	"time"
)

func main() {

	start := time.Now()

	jobs := make(chan bool)
	results := make(chan string)
	done := make(chan bool)

	numberOfWorkers := 10
	numberOfJobs := 1000

	fmt.Println("starting workers")
	for i := 0; i < numberOfWorkers; i++ {
		go worker(jobs, results, done, i) //spawn 10 concurrent worker functions off the main Go routine
	}
	go func() {
		for i := 0; i < numberOfJobs; i++ {
			jobs <- true //keep sending true on the jobs channel 1000 times. will be received by all workers
		}
	}()
	go func() {
		count := 0
		for {
			result := <-results //keep receiving from the results channel until count is >1000
			fmt.Println(result)
			count++
			if count >= numberOfJobs { //1000 responses received
				done <- true //signal Goroutine workers to finish
				return       //exit for loop
			}
		}
	}()
	//main cant continue until there is something to read from the channel so Go routines can so their work
	<-done //receive when numberOfJobs has been completed
	elapsed := time.Since(start)
	fmt.Println("Program %s", elapsed)
}

func worker(jobs chan bool, results chan string, done chan bool, num int) {
	for {
		select {
		case <-jobs: //will keep calling the getResult function until done channel has a message
			res, err := getResult(num)
			if err != nil {
				panic(err)
			}
			results <- res + " Go routine " + strconv.Itoa(num) + "\n"
		case <-done: //when 1000 responses have been received all workers will receive a signal to finish
			return
		}
	}
}

func getResult(num int) (string, error) {
	resp := client.GetResponse("http://localhost:8181/Users/1")
	fmt.Println("\nget request sent from Go routine " + strconv.Itoa(num))
	return resp, nil
}
