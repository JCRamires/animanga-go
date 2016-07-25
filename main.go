package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	populateWorkDetails(report.Works[:])

	session.Close()
}

func populateWorkDetails(works []Work) {
	if len(works) == 0 {
		return
	}
	continueSaving := true
	sliceSize := 50
	if len(works) < 50 {
		sliceSize = len(works)
		continueSaving = false
	}

	saveWorkDetails(works[:sliceSize])

	if continueSaving {
		populateWorkDetails(works[sliceSize:])
	}
}

func saveWorkDetails(works []Work) {
	var url bytes.Buffer
	url.WriteString("http://cdn.animenewsnetwork.com/encyclopedia/api.xml?title=")
	for _, element := range works {
		url.WriteString(strconv.Itoa(element.WorkID))
		url.WriteString("/")
	}

	response, _ := http.Get(url.String())
	fmt.Println(response)
}
