package main

import (
	"fmt"
	//"strconv"
	"strings"
)

func main() {

	portStrings := "80,443,8080,25-30"

	//removeQuotes := strings.Replace(portStrings, "\"", " ", -1)

	var splitString = strings.Split(portStrings, ",")
	fmt.Printf("%q\n", portStrings)

	for _, rune := range splitString {
		if rune == "25-30" { // need to add IF "-" in string
			fmt.Println("TEST")
			rune2 := strings.Split(rune, "-") //[25,30]

			for _, rune := range rune2 {
				// I DONT NEED TO CONVERT TO INT LOL
				//intParsed, err := strconv.Atoi(rune)
				//if err == nil {
				//fmt.Println("%s", intParsed)
				//fmt.Printf("test:%d\n", intParsed)
				fmt.Printf("test:%s\n", rune)

			}
			//fmt.Println(intParsed)
			//fmt.Println(strings.Split(rune, "-"))

		} else {
			fmt.Printf("test:%s\n", rune)
		}
	}

}
