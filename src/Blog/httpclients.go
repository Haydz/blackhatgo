package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// References:
/*
https://golang.org/pkg/io/ioutil/
Package ioutil implements some I/O utility functions.

ReadALL
uses https://golang.org/pkg/io/#Reader a reader


https://golang.org/pkg/os/#FileInfo
from within ReadDir

*/

func main() {

	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Get("http://example.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(body)
	fmt.Printf("%s", body)

	// dirLOL, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for key, value := range dirLOL {
	// 	fmt.Println(key, value.Name(), value.Size())
	// }

	fmt.Println(http.ParseHTTPVersion("HTTP/1.0"))
}
