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
	member      bool            `json:"member"`
	credits     int             `json:"credits"`
	displayName json.RawMessage `json:"display_name"`
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

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main APIKEY term")
	}

	var apiKey = os.Args[1]

	//const BaseURL = "https://api.shodan.io"

	res, err := http.Get(fmt.Sprintf("https://api.shodan.io/account/profile?key=%s", apiKey))
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	//fmt.Println(accountCall)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	//fmt.Println(accountCall.Status)

	s, err := getData([]byte(body))
	fmt.Println(s)

}
