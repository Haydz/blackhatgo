package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

type IpSearch struct {
	RegionCode  json.RawMessage `json:"region_code"`
	IpAddress   string          `json:"ip"`
	History     bool            `json:history(optional)`
	Minimal     bool            `json:minify(optional)`
	CountryName string          `json:"country_name"`
}

//trying account profile information
type AccountProfile struct {
	Member      bool            `json:"member"`
	Credits     int             `json:"credits"`
	DisplayName json.RawMessage `json:"display_name"`
	Created     json.RawMessage `json:"created"`
	//accountCreation string
}

//URL for ip search https://api.shodan.io/shodan/host/134.119.189.29?key={API}
// https://api.shodan.io/shodan/host/{ip}?key={YOUR_API_KEY}

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main APIKEY")
	}

	var apiKey = os.Args[1]
	//var searchTerm = os.Args[2]
	const BaseURL = "https://api.shodan.io"

	introMessage := `[1] - Search a term     [2] - Check Account information [3] - IP Search`

	fmt.Print(introMessage)
	fmt.Println(":")
	reader := bufio.NewReader(os.Stdin)
	selection, _ := reader.ReadString('\n')
	//fmt.Println(reflect.TypeOf(selection))
	selection = strings.TrimSpace(selection)
	selectionInt, err := strconv.Atoi(selection)
	if err != nil {
		fmt.Println(err)
	}

	if selectionInt == 1 {
		fmt.Println("You chose 'Search A Term'")
		//reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the term you wish to search for: ")
		searchTerm, _ := reader.ReadString('\n')
		searchTerm = strings.TrimSpace(searchTerm)
		//fmt.Print("You live in " + city)

		//May try later, idea is to have a function to handle URLs
		// func connect(words ...string){

		//want to add query search now
		res2, err := http.Get(
			fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, apiKey, searchTerm))
		if err != nil {
			fmt.Println(err)
		}
		body2, err := ioutil.ReadAll(res2.Body)
		if err != nil {
			log.Panicln(err)
		}

		var HostTest HostSearch
		_ = json.Unmarshal([]byte(body2), &HostTest)
		for _, host := range HostTest.Matches {
			//fmt.Printf("Host:%18s:%8d\n", strings.TrimSpace(host.IPString), host.Port)
			// add an array so IF HOST OS add

			fmt.Println("Host:", strings.TrimSpace(host.IPString), host.Port, "Hostnames:", host.Hostnames, "OS:", host.OS)
		}

	} else if selectionInt == 2 {
		fmt.Println("You chose Account Information")
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
		fmt.Println(AccountProfileStruct.Member, ":member", AccountProfileStruct.Credits)
		//Attempting with NewDecoder
	} else if selectionInt == 3 {
		fmt.Println("You chose an IP search")
		fmt.Println("Which IP address do you want to search?: ")
		chosenIP, _ := reader.ReadString('\n')
		chosenIP = strings.TrimSpace(chosenIP)
		fmt.Println(chosenIP)

		//Grabbing Account information
		IPRequest, err := http.Get(fmt.Sprintf("%s/shodan/host/134.119.189.29?key=%s&minify=true", BaseURL, apiKey))
		//IPRequest, err := http.Get(fmt.Sprintf("%s/shodan/host/%s?key=%s", BaseURL, chosenIP, apiKey))
		//https://api.shodan.io/shodan/host/{ip}?key={YOUR_API_KEY}
		if err != nil {
			log.Panicln(err)
		}
		defer IPRequest.Body.Close()
		body3, err := ioutil.ReadAll(IPRequest.Body)
		responseJSON3 := string(body3)
		IpSearchStruct := IpSearch{}
		//fmt.Println(body3)
		// //var testAS AccountProfile
		// //reading the json un serializing it
		// // aligning with APIInfo struct
		// //unmarshal because JSON is already in memory
		err3 := json.Unmarshal([]byte(responseJSON3), &IpSearchStruct)
		if err3 != nil {
			panic(err3)
		}
		fmt.Println("==IP Address information")
		fmt.Println(IpSearchStruct.IpAddress)
		//Attempting with NewDecoder
	} else {
		fmt.Println("Not a valid Option")
	}
	//**BROKEN OUT INTO OTHER FILE

}
