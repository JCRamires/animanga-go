package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
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

	var wg sync.WaitGroup
	populateWorkDetails(report.Works[:], &wg)
	wg.Wait()
}

func populateWorkDetails(works []Work, wg *sync.WaitGroup) {
	if len(works) == 0 {
		return
	}
	continueSaving := true
	sliceSize := 50
	if len(works) < 50 {
		sliceSize = len(works)
		continueSaving = false
	}

	go getDetailsWorker(works[:sliceSize], wg)

	time.Sleep(time.Second)

	if continueSaving {
		populateWorkDetails(works[sliceSize:], wg)
	}
}

func getDetailsWorker(works []Work, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var url bytes.Buffer
	url.WriteString("http://cdn.animenewsnetwork.com/encyclopedia/api.xml?title=")
	for _, element := range works {
		url.WriteString(strconv.Itoa(element.WorkID))
		url.WriteString("/")
	}

	response, _ := http.Get(url.String())

	xmlFile, _ := ioutil.ReadAll(response.Body)

	var detailsReport DetailsReport
	xml.Unmarshal(xmlFile, &detailsReport)

	session, err := getDatabaseSession()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	collection := getDetailsCollection(session)
	for _, element := range detailsReport.Anime {
		collection.Insert(element)
	}

	// for _, element := range detailsReport.Manga {
	// 	collection.Insert(element)
	// }

	response.Body.Close()
}
