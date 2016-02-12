package database

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Mongodb struct {
	session *mgo.Session
}

func GetSession() (session *mgo.Session) {

	if session == nil {
		var err error
		session, err = mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
	}
	session.SetMode(mgo.Monotonic, true)
	return session.Copy()
}
