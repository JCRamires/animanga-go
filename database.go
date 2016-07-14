package main

import "gopkg.in/mgo.v2"

func getDatabaseSession() (session *mgo.Session, err error) {
	session, err = mgo.Dial("localhost:27017")
	return
}

func getWorksCollection(session *mgo.Session) (collection *mgo.Collection) {
	collection = session.DB("animanga-go").C("works")

	index := mgo.Index{
		Key:        []string{"workid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	collection.EnsureIndex(index)

	return
}
