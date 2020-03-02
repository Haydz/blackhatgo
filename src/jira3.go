package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatalln("Usage: main user pass")
	}
	var userName = os.Args[1]
	var password string = "test"
	client := &http.Client{}
	//adding proxy authentication
	auth := strings.TrimSpace(userName + ":" + password)
	fmt.Println(auth)
	request, err := http.NewRequest("GET", "https://jira.points.com/rest/api/latest/issue/SEC-375", nil)
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Add("Authorization", basicAuth)

	//printing the request to the console
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	//calling the URL
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	log.Println(response.StatusCode)
	log.Println(response.Status)
}

// func basicAuth(userName string, password string) string {

// 	//http://localhost:8080/rest/api/2/issue/createmeta
// 	//https://jira.atlassian.com/rest/api/latest/issue/JRA-9

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", "https://jira.points.com/rest/api/latest/issue/SEC-375", nil)
// 	req.SetBasicAuth(userName, password)
// 	//fmt.Println(req.Status)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bodyText, err := ioutil.ReadAll(resp.Body)
// 	s := string(bodyText)
// 	return s
// }
