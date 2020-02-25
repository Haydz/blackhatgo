package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	respGet, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}

	//Print HTTP Status
	fmt.Println(respGet.Status, "GET")
	//read and display response body
	body, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	//close connection
	respGet.Body.Close()

	respHead, err := http.Head("https://google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	respHead.Body.Close()
	respHead.Body.Close()
	fmt.Println(respHead.Status, "HEAD")

	//defer respHead.Body.Close()

	//post request
	form := url.Values{}
	form.Add("foo", "bar")
	respPost, err := http.PostForm("https://google.com/robots.txt", form)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(respPost.Status, "POST")

	//other VERBS with client
	var client http.Client
	req, err := http.NewRequest("PUT",
		"https://google.com/robots.txt", strings.NewReader(form.Encode()))
	resp, err := client.Do(req)
	resp.Body.Close()
	fmt.Println(resp.Status, "PUT WITH NEWREQUEST")
}
