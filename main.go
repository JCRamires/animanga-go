package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// response, _ := http.Get("http://www.animenewsnetwork.com/encyclopedia/reports.xml?id=155&nlist=all")
	response, _ := http.Get("http://www.animenewsnetwork.com/encyclopedia/reports.xml?id=155")

	xmlFile, _ := ioutil.ReadAll(response.Body)

	response.Body.Close()

	var report Report
	xml.Unmarshal(xmlFile, &report)

	session, err := getDatabaseSession()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	collection := getWorksCollection(session)
	for _, element := range report.Works {
		collection.Insert(element)
	}

	session.Close()
}
