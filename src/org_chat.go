package main

import "fmt"

type myStruct struct {
	name     string
	manager  string
	position string
}

type mapKey struct {
	key string
}

func main() {

	testStruct := make(map[mapKey]myStruct)

	testStruct[mapKey{"Roger"}] = myStruct{manager: "Tyler", position: "Senior Development Engineer"}
	testStruct[mapKey{"Tyler"}] = myStruct{manager: "Trey", position: "Engineering Manager"}
	testStruct[mapKey{"Trey"}] = myStruct{position: "CTO"}
	fmt.Println(testStruct)

	for key, _ := range testStruct {
		fmt.Println(key)
	}

	fmt.Println(testStruct.key["Roger"].position)
	// mapOfMaps := make(map[string]map[string]myStruct)
	// mapOfMaps[name] = make(map[string]myStruct)

}

// 	if _, ok := mapOfMaps[key]; !ok {
// 		mapOfMaps[key] = make(map[string]myStruct)
// 	}

// 	if _, ok := mapOfMaps[key][key2]; !ok {
// 		mapOfMaps[key][key2] = new(myStruct)
// 	}

// }
