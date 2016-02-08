package server

import (
	"fmt"
	"net/http"
)

func Startserver() {
	fmt.Println("Server running")
	err := http.ListenAndServe(":8181", nil)
	checkErrorAndPanic(err)
}

func checkErrorAndPanic(err error) {

	if err != nil {
		panic(err)
	}

}
