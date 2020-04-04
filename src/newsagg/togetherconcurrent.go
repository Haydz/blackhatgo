package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsRoutine(c chan News, Location string) {
	var n News
	defer func() { c <- n }()

	resp, err := http.Get(Location)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	//c <- n

}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	var s Sitemapindex
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	news_map := make(map[string]NewsMap)
	resp.Body.Close()
	queue := make(chan News, 40)

	for _, Location := range s.Locations {
		Location = strings.TrimSpace(Location)
		wg.Add(1)
		// 	fmt.Println("FOUND SITE")
		// 	fmt.Println(Location)
		newsRoutine(queue, Location)
	}
	wg.Wait()
	close(queue)

	for elem := range queue {
		for idx := range elem.Keywords {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}

	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}

	t, _ := template.ParseFiles("newsaggtemplate.html")
	t.Execute(w, p)

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
