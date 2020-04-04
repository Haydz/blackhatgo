package main

import (
	"fmt"
)

type myStruct struct {
	manager  string
	position string
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	var myMap = map[string]myStruct{

		"Tyler": {manager: "Trey", position: "Engineering Manager"},
		"Roger": {manager: "Tyler", position: "Senior Development Engineer"},
		"Trey":  {position: "CTO"},
		"Chris": {position: "CEO"},
		"Mike":  {manager: "Chris", position: "Hiring Specialist"},
		"Anna":  {manager: "Trey", position: "Senior Softeware Dev"},
		"John":  {manager: "Tyler", position: "Development Engineer"},
		"Alfed": {manager: "Roger", position: "Junior Engineer"},
		"Ralph": {manager: "Chris", position: "Finances Expert"},
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
		//fmt.Println("==Searching to see who Clevel Manages: ", cLevelNames[i])
		//Search through map for who C level managers
		println("##", cLevelNames[i])
		for key, value := range myMap {

			if value.manager == cLevelNames[i] {

				//fmt.Println(cLevelNames[i], "manages:", key)
				manager = append(manager, key)
				indivudalManager := key
				//looking to see if Manager manages anyone
				println("####", indivudalManager)

				for key, value := range myMap {
					if value.manager == indivudalManager {
						//fmt.Println(indivudalManager, "has an employee", key)
						println("######", key)
					}

				}

			}
		}

		//find sub managers
		//fmt.Println("==middle/ Clevel employees manager anyone??==")

		//fmt.Println(cLevelNames[i], "manager array:", manager)
	}

	// find all managers
	fmt.Println("===========================$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	var allManagers []string
	manager2 := ""
	for _, value := range myMap {
		manager2 = value.manager
		if len(allManagers) == 0 {
			//manager2 = value.manager
			//fmt.Println(manager2)
			allManagers = append(allManagers, manager2)
			fmt.Println("Adding initial valuel", allManagers)
		} else {
			fmt.Println("else loop using: ", value.manager)
			// nee dto check IF manager is in Allmanagers
			compare := Contains(allManagers, manager2)
			if compare != true {
				allManagers = append(allManagers, manager2)
			}
			// for i := 0; i < len(allManagers); i++ {
			// 	if manager2 == allManagers[i] {
			// 		//manager2 = value.manager
			// 		continue
			// 	}
			// }

		}
	}
	fmt.Println("Allmanagers array: ", allManagers)

}

//fmt.Println("ALL managers array:", allManagers)
// for i := 0; i < len(allManagers); i++ {
// 	if manager2 == allManagers[i] {
// 		//manager2 = value.manager
// 		continue
// 	}
// }
