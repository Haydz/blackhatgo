package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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

//trying account profile information
type AccountProfile struct {
	Member      bool            `json:"member"`
	Credits     int             `json:"credits"`
	DisplayName json.RawMessage `json:"display_name"`
	Created     json.RawMessage `json:"created"`
	//accountCreation string
}

func main() {

	if len(os.Args) != 3 {
		log.Fatalln("Usage: main APIKEY term")
	}

	var apiKey = os.Args[1]
	var searchTerm = os.Args[2]
	const BaseURL = "https://api.shodan.io"

	//May try later, idea is to have a function to handle URLs
	// func connect(words ...string){

	// 	testConn, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, apiKey))
	// }

	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, apiKey))
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}

	responseJSON := string(body)
	testStruct := APIInfo{}
	//reading the json un serializing it
	// aligning with APIInfo struct
	//unmarshal because JSON is already in memory
	_ = json.Unmarshal([]byte(responseJSON), &testStruct)
	fmt.Println("== YOUR API Information")
	fmt.Println(testStruct.QueryCredits, ":credits\n",
		testStruct.Plan, ":plan")

	//want to add query search now
	res2, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, apiKey, searchTerm))
	//"https://api.shodan.io/shodan/host/search?key=T2VqWX5fJWIXpZ1I5tMXRMVhAKRqqWbA&query=hp")
	if err != nil {
		fmt.Println(err)
	}
	//Parsing JSON from HOST Search
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

	//**BROKEN OUT INTO OTHER FILE
	//Grabbing Account information
	AccountCall, err := http.Get(fmt.Sprintf("%s/account/profile?key=%s", BaseURL, apiKey))
	//res3, err := http.Get(fmt.Sprintf("https://api.shodan.io/account/profile?key="))
	if err != nil {
		log.Panicln(err)
	}
	defer AccountCall.Body.Close()
	// //fmt.Println(accountCall)
	body3, err := ioutil.ReadAll(AccountCall.Body)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// //fmt.Println(accountCall.Status)

	responseJSON3 := string(body3)
	AccountProfileStruct := AccountProfile{}
	// //var testAS AccountProfile
	// //reading the json un serializing it
	// // aligning with APIInfo struct
	// //unmarshal because JSON is already in memory
	err3 := json.Unmarshal([]byte(responseJSON3), &AccountProfileStruct)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("== YOUR Account Information")
	fmt.Println(AccountProfileStruct.Member, ":member")
	//Attempting with NewDecoder
	var ret APIInfo
	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, apiKey))
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		panic(err3)
	}
	fmt.Println(ret.Telnet, "Telnet", ret.ScanCredits, "Scan Credits")
}
