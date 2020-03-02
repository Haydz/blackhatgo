package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

// type Location struct {
// 	Loc string `xml:"loc"`
// }

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func main() {
	var s Sitemapindex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()
	//fmt.Println(s.Locations)
	news_map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		//fmt.Printf("%s\n", Location)
		//fmt.Println(Location, "TEST")
		Location = strings.TrimSpace(Location)
		fmt.Println(isValidUrl(Location))

		fmt.Println("IS A VALID URL")
		resp2, err := http.Get(Location)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status)
		bytes, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			fmt.Println(err)
		}
		xml.Unmarshal(bytes, &n)
		//fmt.Println(n.Titles, "\n", n.Locations)
		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}

}
