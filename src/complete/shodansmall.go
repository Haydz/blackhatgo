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

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main APIKEY term")
	}

	var apiKey = os.Args[1]

	const baseURL = "https://api.shodan.io"

	res, _ := http.Get(fmt.Sprintf("%s/account/profile?key=%s", baseURL, apiKey))
	Profile := AccountProfile{}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	_ = json.Unmarshal([]byte(string(body)), &Profile)

	fmt.Println(Profile.Credits)

	//OTHER WAY
	//with decoder can create struct in 2 ways
	//Profile2 := AccountProfile{}
	var Profile2 AccountProfile
	try2, err := http.Get(fmt.Sprintf("%s/account/profile?key=%s", baseURL, apiKey))
	err2 := json.NewDecoder(try2.Body).Decode(&Profile2)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(Profile2.Credits)

}
