package main

import (
	"fmt"
	"strings"
)

func main() {

	portStrings := "80,443,8080"

	//removeQuotes := strings.Replace(portStrings, "\"", " ", -1)

	var splitString = strings.Split(portStrings, ",")
	fmt.Printf("%q\n", portStrings)

	for _, rune := range splitString {
		fmt.Printf("test:%s80\n", rune)
	}

}
