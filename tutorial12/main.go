package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex stores a slice of Locations that we decode from an XML sitemap
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location stores a single location decoded from an XML sitemap
type Location struct {
	Loc string `xml:"loc"`
}

/*
	Converts each location to a string (from an object)
	from {www.url.com/target} ==> "www.url.com/target"
*/
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

/*
	Tutoial 12 we learned:
		* looping with range (over structs)
		* for {} is an endless loop until we call 'break'
		* for i := 0 ; i < 10 ; i++ {} is basic format for looping
		* for i:= 0 ; a <10 ; i++ {} we can customize our conditions
		* bottom line: efficient and customizable
*/
func main() {
	response, getErr := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if getErr != nil {
		fmt.Println(getErr)
	}
	bytes, ioErr := ioutil.ReadAll(response.Body)
	if ioErr != nil {
		fmt.Println(ioErr)
	}
	response.Body.Close()
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// unused value is index
	for _, Location := range s.Locations {
		fmt.Printf("%s", Location)
	}
}
