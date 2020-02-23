package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	respGET, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}

	//print http status
	fmt.Println(respGET.Status)

	//Read and display response body

	body, err := ioutil.ReadAll(respGET.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(body))
	respGET.Body.Close()

}
