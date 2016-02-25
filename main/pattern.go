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
		go worker(jobs, results, done, i)
	}
	go func() {
		for i := 0; i < numberOfJobs; i++ {
			jobs <- true
		}
	}()
	go func() {
		count := 0
		for {
			result := <-results
			fmt.Println(result)
			count++
			if count >= numberOfJobs {
				done <- true
				return
			}
		}
	}()
	<-done
	elapsed := time.Since(start)
	fmt.Println("Program %s", elapsed)
}

func worker(jobs chan bool, results chan string, done chan bool, num int) {
	for {
		select {
		case <-jobs:
			res, err := getResult(num)
			if err != nil {
				panic(err)
			}
			results <- res + " Go routine " + strconv.Itoa(num) + "\n"
		case <-done:
			return
		}
	}
}

func getResult(num int) (string, error) {
	resp := client.GetResponse("http://localhost:8181/Users/1")
	fmt.Println("\nget request sent from Go routine " + strconv.Itoa(num))
	return resp, nil
}
