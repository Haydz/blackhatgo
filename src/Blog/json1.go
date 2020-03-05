package main

import (
	"encoding/json"
	"fmt"
)

//Person - VSCODE forces me to comment above a struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"Age"`
}

func main() {

	//Create instance of Person
	PersonA := Person{Name: "John", Age: 14}

	fmt.Println("PersonA format tied to Person struct prior to Marshalling")
	// Struct before marshalling

	fmt.Printf("%+v\n", PersonA)

	//Marshall JSON
	a, err := json.Marshal(PersonA)
	if err != nil {
		fmt.Println("Marshal didn't work")
	}

	fmt.Println("Marshalled PersonA in Binary Format")
	fmt.Println(a)

	fmt.Println("Marshalled PersonA in String format")
	fmt.Println(string(a))

	fmt.Println("test")

}
