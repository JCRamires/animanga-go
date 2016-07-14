package main

// Report struct
type Report struct {
	Works []Work `xml:"item"`
}

// Work struct
type Work struct {
	GUID      string `xml:"gid"`
	Title     string `xml:"name"`
	Precision string `xml:"precision"`
}
