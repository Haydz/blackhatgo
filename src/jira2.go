package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatalln("Usage: main user pass")
	}
	var userName = os.Args[1]
	var password = os.Args[2]
	a := (basicAuth("userName", "password"))
	fmt.Println(a)
}

func basicAuth(userName string, password string) string {

	//http://localhost:8080/rest/api/2/issue/createmeta
	//https://jira.atlassian.com/rest/api/latest/issue/JRA-9

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://jira.points.com/rest/api/latest/issue/SEC-375", nil)
	req.SetBasicAuth(userName, password)
	//fmt.Println(req.Status)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}
