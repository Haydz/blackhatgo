package main

import (
	"fmt"
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
		"Anna":  {manager: "Trey", position: "Senior Softeware Dev"},
	}

	// I need to create a map that ties

	var cLevelNames []string

	fmt.Println("===trying in while loop")

	// loop to find all top levels of trees( clevels)

	for key, value := range myMap {
		//find all C levels and place into cLevelNames

		if value.manager == "" {
			fmt.Println("C level found", key)
			cLevelNames = append(cLevelNames, key)

		}

	}

	//fmt.Println("cLevel Slice:", cLevelNames)

	//loop to crawl through cLevels found
	fmt.Println("==loop for finding who C levels manage (C  level Employees)")
	for i := 0; i < len(cLevelNames); i++ {
		// manager array here so it wipes clean each C level:
		var manager []string

		fmt.Println("Searching through Clevel tree:", cLevelNames[i])
		//Search through map for who C level managers
		for key, value := range myMap {
			if value.manager == cLevelNames[i] {
				fmt.Println("==Searching to see who Clevel Manages: ", cLevelNames[i])
				fmt.Println(cLevelNames[i], "manages:", key)
				manager = append(manager, key)
			}
		}
		//find sub managers
		fmt.Println("==middle/ Clevel employees manager anyone??==")
		for i := 0; i < len(manager); i++ {

			for key, value := range myMap {
				if value.manager == manager[i] {
					fmt.Println(manager[i], "has an employee", key)
				}

			}
		}
		fmt.Println(cLevelNames[i], "manager array:", manager)
	}

}

//find managers

// 	b := 0 // loop through whole map
// 	for b < len(myMap) {
// 		name2 := ""
// 		for key := range myMap {
// 			string := "#"
// 			if cLevel == false && name2 == "" {
// 				//fmt.Println(strings.Repeat(string, i2), key)
// 				fmt.Println("found a C-level - top of forest)")
// 				name2 = key
// 				b++
// 				//fmt.Println(value)
// 				//cLevel = true
// 				for key, value := range myMap {
// 					if name2 == key && value.manager == "" { //we found a CEO //

// 						i2++
// 						fmt.Println(strings.Repeat(string, i2), key)
// 						reader := bufio.NewReader(os.Stdin)
// 						selection, _ := reader.ReadString('\n')
// 						fmt.Println(selection)

// 						cLevel = true
// 						for key, value := range myMap {
// 							if name2 == value.manager { // if has manager, must be employee}
// 								fmt.Println("manager section")
// 								i2++
// 								fmt.Println(strings.Repeat(string, i2), key)

// 							}
// 						}
// 					} // found base of tree.
// 				}
// 			} else if cLevel == true {
// 				i2 = 0
// 				cLevel = false
// 				name2 = ""
// 				continue
// 			}
// 		}
// 	}
// }
