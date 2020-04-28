package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Expand struct {
	Expand     string `json:"expand"`
	MaxResults int    `json:maxResults`
	Total      int    `json:total`
	//Issues     []Issue
}
type Issues struct {
	Issues []Issue
}

type Issue struct {
	Id     string `json:id`
	Key    string `json:key`
	Self   string `json:self`
	Fields Fields
	Names  Names
}

type Fields struct {
	Summary  string `json:summary`
	Priority Priority
	Assignee Assignee
	Created  SpecialDate `json:created`
	Updated  SpecialDate `json:updated`
	Status   Status
}

type Status struct {
	Name string `json:name`
}

type SpecialDate struct {
	time.Time
}

type Priority struct {
	Self string `json:self`
	Name string `json:name`
}

type Assignee struct {
	Self string `json:self`
	Name string `json:name`
}

type Names struct {
	AdditionalProperties string `json:Additional Properties`
}

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02T15:04:05.999999999Z0700", strInput)
	if err != nil {
		return err
	}

	sd.Time = newTime
	return nil
}

/*
References
http://polyglot.ninja/golang-making-http-requests/
https://stackoverflow.com/questions/19253469/make-a-url-encoded-post-request-using-http-newrequest
https://appdividend.com/2019/12/02/golang-http-example-get-post-http-requests-in-golang/
https://girishjoshi.io/post/implementing-http-basic-authentication-in-golang/
https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/

Jira:
https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-rest-api-3-filter-defaultShareScope-put
https://godoc.org/github.com/andygrunwald/go-jira#FilterSearchOptions
https://godoc.org/github.com/andygrunwald/go-jira#pkg-examples

https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-rest-54: Search for issues using JQL (GET)

https://tanaikech.github.io/2017/09/15/spreadsheets.values.batchupdate-using-golang/

https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-rest-api-3-search-get

*/

func queryJira(username string, password string, query []string, dataToAdd []string) []byte {

	username = username
	passwd := password

	//function to call Jira, allow different queries
	apiURL := "https://jira.points.com/rest/api/2/search"
	resource := "/rest/api/2/search"
	data := url.Values{}
	data.Set(query[0], query[1])

	for _, value := range dataToAdd {
		data.Add("fields", value)
	}

	//building URL
	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)
	fmt.Println(urlStr)

	client2 := &http.Client{}
	r, _ := http.NewRequest("GET", urlStr, nil)
	//setting basicAuth credentials
	r.SetBasicAuth(username, passwd)
	//setting head for JSON
	r.Header.Add("Accept", "application/json")
	resp, err := client2.Do(r)
	if err != nil {
		fmt.Println("error connecting: ", err)
	}
	defer resp.Body.Close()

	//reading body response
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading data after being sent", err)
	}

	return bodyText

}

func main() {

	//== THIS IS FOR ENTERING CREDENTIALS ON THE COMMAND LINE
	// UNCOMMENT WHEN CODE IS FINISHED ========
	// if len(os.Args) != 3 {
	// 	log.Fatalln("Usage: main username password")
	// }
	// var username = os.Args[1]
	// var password = os.Args[2]
	username := "NULL"
	password := "NULL"
	//queries and fields wanted

	//fmt.Println(jiraResponse)
	// fmt.Println(string(jiraResponse))
	// var expando Expand
	// json.Unmarshal([]byte(jiraResponse), &expando)
	// fmt.Println("Max Results: ", expando.MaxResults, " Number of Issues: ", expando.Total)

	//JIRA Options;
	// Spreadsheet ID

	//Fields to grab
	dataToAdd := []string{"key", "updated", "summary", "created", "status", "priority", "assignee"}

	//Organizing SRR to write
	query := []string{"jql", "project=srr and status != Resolved  AND  issuetype not in subtaskIssueTypes()"}
	jiraResponse := queryJira(username, password, query, dataToAdd)

	var sRROpen Issues
	json.Unmarshal([]byte(jiraResponse), &sRROpen)
	//fmt.Println(issues)
	for _, issue := range sRROpen.Issues {
		fmt.Println(issue.Key)
		fmt.Println(issue.Self)
		fmt.Println(issue.Fields.Summary)
		fmt.Println(issue.Fields.Priority.Name)
	}

	//Organizing Open Pentest write
	pentestQuery := []string{"jql", "labels = pentesting   AND type != Sub-task AND resolution = unresolved AND project = SEC ORDER BY createdDate"}
	pentestResponse := queryJira(username, password, pentestQuery, dataToAdd)
	var pentestOpen Issues
	json.Unmarshal([]byte(pentestResponse), &pentestOpen)
	//fmt.Println(pentestIssues)
	fmt.Println("\n\n\n")
	fmt.Println("===PENTEST ISSUES HERE=====")
	for _, issue := range pentestOpen.Issues {
		fmt.Println(issue.Key)
		fmt.Println(issue.Self)
		fmt.Println(issue.Fields.Summary)
		fmt.Println(issue.Fields.Priority.Name)
		fmt.Println(issue.Fields.Assignee.Name)
		fmt.Println(issue.Fields.Updated)
		fmt.Println(issue.Fields.Status.Name)
	}
}

//https://www.prudentdevs.club/gsheets-go
