package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	return nil
}

func main() {

	if len(os.Args) != 3 {
		log.Fatalln("Usage: main user pass")

		//http://localhost:8080/rest/api/2/issue/createmeta
		//https://jira.atlassian.com/rest/api/latest/issue/JRA-9
	}
	var userName = os.Args[1]
	var password = os.Args[2]

	// client := &http.Client{
	// 	Jar:           cookieJar,
	// 	CheckRedirect: redirectPolicyFunc,
	// }

	req, err := http.NewRequest("GET", "https://jira.points.com/rest/api/latest/issue/SEC-375", nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(userName, password))

	//resp, err := client.Do(req)
}
