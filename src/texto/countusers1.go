package main

import "fmt"

func main() {
	//Initialize an array
	inputArray := []string{"Tim", "Haydn", "John", "Haydn", "Haydn"}
	//listUsers := []string{"Tim", "Haydn", "John", "Haydn"}
	printUniqueValue(inputArray)
}

func printUniqueValue(listusers []string) {
	//Create a   dictionary of values for each element
	dict := make(map[string]int)
	for _, name := range listusers {
		//fmt.Println(key)
		//dict[name] is adding the name to the MAP key
		// when you want the value for the KEY (name) of a map, you use the syntax dict[NAME]
		// hency dict[name] +1 actually means the value of dictionary key (name) + 1
		dict[name] = dict[name] + 1
		//fmt.Println(name)
		//fmt.Println(dict[name])
	}
	fmt.Println(dict)
	fmt.Println(dict["Haydn"])
}
