package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//open file
	file, err := os.Open("sshinvalid.txt")
	if err != nil {
		fmt.Println("Opening file error", err)
	}
	//need to scan the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var words []string // so I can play with it easier and outside of the scan loop

	for scanner.Scan() {
		words = append(words, scanner.Text())

	}

	//fmt.Println(words)

	/* want to seperate words by space
	Use the fields function

	*/
	for _, value := range words {
		wordBreakDown := strings.Fields(value)
		fmt.Println("==breaking sentence into words==")
		for _, value := range wordBreakDown {
			fmt.Println(value)
		}

	}

	// want to seperate via something else - lets use a comma
	separateInput := "This,is,a,CSV,Comma,separated,Words"
	separatingString := strings.FieldsFunc(separateInput, func(r rune) bool {
		if r == ',' {
			return true
		}
		return false
	})

	fmt.Println("===This is the string separated by comma===")
	for i, value := range separatingString {
		fmt.Println(i, ":", value)
	}

	// lets see if string HAS something, similar to if IN:
	for _, value := range words {
		if strings.Contains(value, "logname=") == true {
			fmt.Println("Found sentence with logname:::", value)
		}
	}

	//now lets separate via = sign and attempted to get the IP address + user
	for _, value := range words {
		if strings.Contains(value, "logname=") == true {
			fmt.Println("Found sentence with logname:::", value)
		}
	}

	//separate via TWO delimiters, similar to a cut
	// lets grab the IP address
	/*
	    cat sshinvalid.txt  | grep logname | cut -d "[" -f2 | cut -d "]" -f1
	   5153
	   5153
	*/
	var splitPorts []string
	for _, value := range words {
		if strings.Contains(value, "logname=") == true {
			// fmt.Println("Found sentence with logname:::", value)
			cuttingByTwo := strings.FieldsFunc(value, func(r rune) bool {
				if r == '[' || r == ']' {
					return true
				}
				return false
			})
			splitPorts = append(splitPorts, cuttingByTwo[1])
		}

	}

	fmt.Println("Ports found split buy [ & ]: ", splitPorts)
}

//want to sort
