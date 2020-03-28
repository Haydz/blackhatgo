package main

import (
	"fmt"
	"strings"
)

type myStruct struct {
	manager  string
	position string
}

func main() {
	var myMap = map[string]myStruct{

		"Tyler": {manager: "Trey", position: "Engineering Manager"},
		"Roger": {manager: "Tyler", position: "Senior Development Engineer"},
		"Trey":  {position: "CTO"},
		"Chris": {position: "CEO"},
		"Mike":  {manager: "Chris", position: "Hiring Specialist"},
	}

	fmt.Println("Rogers Manager: ", myMap["Roger"].manager)
	fmt.Println("Tylers Manager :", myMap["Tyler"].manager)
	fmt.Println("Treys Manager :", myMap["Trey"].manager)
	//fmt.Println(myMap["Roger"].position)

	if myMap["Trey"].manager == "" {
		fmt.Println("empty")
	}

	// if no manager then goes on top because boss

	// fmt.Println("==== TRYING LOOP")
	// name := ""
	// i := 0
	// for key, value := range myMap {

	// 	string := "#"

	// 	if value.manager == "" {
	// 		i++
	// 		fmt.Println(strings.Repeat(string, i), key)
	// 		name = key
	// 	} else if value.manager == name {
	// 		i++
	// 		fmt.Println(strings.Repeat(string, i), key)
	// 		name = key
	// 	}

	// }

	fmt.Println("===trying in while loop")

	//fmt.Println(len(myMap))
	i2 := 1 // string #
	//c := 0  // C levels found

	//https://play.golang.org/p/4w9C79woV5-
	cLevel := False
	b := 0  // loop through whole len
	for b < len(myMap) {
		name2 := ""
		for key, value := range myMap {
			string := "#"
			// start of loop has begin, b = 0
			// we want to find person with no manager
			if (cLevel == False && name2 =="") {
					i2++
					fmt.Println(strings.Repeat(string, i2), key)
					name2 = key
					b++
				} else {
					b++
					continue
				}
			} else if c != 0 && name2 == "" {
				i2 = 0
				fmt.Println(strings.Repeat(string, i2), key)
				name2 = key
				b++

			} else if value.manager == name2 {
				i2++
				fmt.Println(strings.Repeat(string, i2), key)

				b++
				name2 = key

			} else {
			}
		}
	}
}
