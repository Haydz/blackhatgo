package main

import "fmt"

func main() {
	//var grades map[string]float32

	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["jess"] = 92
	grades["sam"] = 71

	fmt.Println(grades)

	timsGrade := grades["Timmy"]
	fmt.Println(timsGrade)

	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}
