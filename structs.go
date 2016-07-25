package main

// Report struct
type Report struct {
	Works []Work `xml:"item"`
}

// DetailsReport struct
type DetailsReport struct {
	Anime []Anime `xml:"anime"`
	Manga []Manga `xml:"manga"`
}

// Work struct
type Work struct {
	WorkID    int    `xml:"id"`
	Title     string `xml:"name"`
	Precision string `xml:"precision"`
}

// Anime struct
type Anime struct {
	WorkID    int    `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Precision string `xml:"precision,attr"`
}

// Manga struct
type Manga struct {
}
