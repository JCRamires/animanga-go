package main

// Report struct
type Report struct {
	Works []Work `xml:"item"`
}

// Work struct
type Work struct {
	WorkID    int    `xml:"id"`
	Title     string `xml:"name"`
	Precision string `xml:"precision"`
}
