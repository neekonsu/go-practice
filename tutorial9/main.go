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
	From tutorials (9 -> 11):
		We learned to:
			* Make GET requests
			* Make structs for decoding XML content
			* Convert from a struct to a primitive type
				=> from ('Location' ==> 'string')
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
	fmt.Println(s.Locations)
}
