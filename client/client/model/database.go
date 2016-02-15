package model

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)

func ConnectToMongo() bool {
	ret := false
	fmt.Println("enter main - connecting to mongo")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Detected panic")
			var ok bool
			err, ok := r.(error)
			if !ok {
				fmt.Printf("pkg:  %v,  error: %s", r, err)
			}
		}
	}()

	maxWait := time.Duration(5 * time.Second)
	session, sessionErr := mgo.DialWithTimeout("localhost:27017", maxWait)
	if sessionErr == nil {
		session.SetMode(mgo.Monotonic, true)
		coll := session.DB("MyDB").C("MyCollection")
		if coll != nil {
			fmt.Println("Got a collection object")
			ret = true
		}
	} else {
		fmt.Println("Unable to connect to local mongo instance!")
	}
	return ret
}
