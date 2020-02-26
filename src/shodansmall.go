package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AccountProfile struct {
	Member      bool            `json:"member"`
	Credits     int             `json:"credits"`
	DisplayName json.RawMessage `json:"display_name"`
	//accountCreation string
}

func getData(body []byte) (*AccountProfile, error) {
	var s = new(AccountProfile)
	err2 := json.Unmarshal(body, &s)
	if err2 != nil {
		fmt.Println("whoops:", err2)
	}
	return s, err2
}

func main() {

	Profile := AccountProfile{}

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main APIKEY term")
	}

	var apiKey = os.Args[1]

	const baseURL = "https://api.shodan.io"

	res, _ := http.Get(fmt.Sprintf("%s/account/profile?key=%s", baseURL, apiKey))
	err2 := json.NewDecoder(res.Body).Decode(&Profile)
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	_ = json.Unmarshal([]byte(string(body)), &Profile)

	fmt.Println(Profile.Credits)
	// 	if err != nil {
	// 		log.Panicln(err)
	// 	}
	// 	defer res.Body.Close()
	// 	//fmt.Println(accountCall)
	// 	fmt.Println(res.Status)
	// 	body, err := ioutil.ReadAll(res.Body)
	// 	if err != nil {
	// 		log.Panicln(err)
	// 	}
	// 	fmt.Println(body)
	// 	//fmt.Println(accountCall.Status)

	// 	s, err := getData([]byte(body))
	// 	fmt.Println(s)

}
