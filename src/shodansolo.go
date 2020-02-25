package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//const BaseURL = "https://api.shodan.io"

//const apiKey = "T2VqWX5fJWIXpZ1I5tMXRMVhAKRqqWbA"

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

type HostLocation struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float32 `json:"latitude"`
}

type Host struct {
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	ISP       string       `json:"isp"`
	ASN       string       `json:"asn"`
	Hostnames []string     `json:"hostnames"`
	Location  HostLocation `json:"location"`
	IP        int64        `json:"ip"`
	Domains   []string     `json:"domains"`
	Org       string       `json:"org"`
	Data      string       `json:"data"`
	Port      int          `json:"port"`
	IPString  string       `json:"ip_str"`
}

type HostSearch struct {
	Matches []Host `json:"matches"`
}

func main() {

	if len(os.Args) != 3 {
		log.Fatalln("Usage: main APIKEY term")
	}

	var apiKey = os.Args[1]
	var searchTerm = os.Args[2]
	const BaseURL = "https://api.shodan.io"
	//res, err := http.Get("https://api.shodan.io/api-info?key=T2VqWX5fJWIXpZ1I5tMXRMVhAKRqqWbA")
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, apiKey))
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		fmt.Println(err)

	}
	fmt.Println(&ret.QueryCredits)
	//.Println("TEST DECODER")
	// fmt.Printf(
	// 	"Query Credits: %d\nScan Credits:  %d\n\n",
	// 	testStruct2.QueryCredits,
	// 	testStruct2.ScanCredits)

	responseJson := string(body)
	testStruct := APIInfo{}
	//reading the json un serializing it
	// aligning with APIInfo struct
	//unmarshal because JSON is already in memory
	_ = json.Unmarshal([]byte(responseJson), &testStruct)
	fmt.Println(testStruct.QueryCredits, "unmarshell credits", "\n", testStruct.Plan, "plan")

	//var requestedReturn APIInfo

	//tested := json.NewDecoder(res.Body).Decode(&requestedReturn)
	//fmt.Println(tested)
	//fmt.Println(json.NewDecoder(res.Body).Decode(&requestedReturn))
	//fmt.Println(&tested, "TEST")
	//info, err := tested.APIInfo()

	// fmt.Printf(
	// 	"Query Credits: %d\nScan Credits:  %d\n\n",
	// 	tested.QueryCredits,
	// 	tested.ScanCredits)

	//want to add query search now
	res2, err := http.Get(
		//fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, apiKey, searchTerm),
		"https://api.shodan.io/shodan/host/search?key=T2VqWX5fJWIXpZ1I5tMXRMVhAKRqqWbA&query=hp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(searchTerm, apiKey)
	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		log.Panicln(err)
	}
	//fmt.Println(res2.Status)

	var HostTest HostSearch
	_ = json.Unmarshal([]byte(body2), &HostTest)
	for _, host := range HostTest.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
