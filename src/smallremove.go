package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func main() {
	//slice to show options chosen
	var dataToAdd []string
	//creating reader for reading in input
	reader := bufio.NewReader(os.Stdin)
	//fields to select
	fieldsToChoose := []string{"created", "updated", "summary", "status", "priority", "assignee"}
	//unending loop until "ENTER" is pressed.
	for x := 0; x < 2; {
		fmt.Println("Please Enter Fields to print out")
		fmt.Println(fieldsToChoose)

		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)

		if value == "" {
			break
		}
		//remove chosen field from options list
		fieldsToChoose = remove(fieldsToChoose, Find(fieldsToChoose, value))
		//append option to a list of chosen options
		dataToAdd = append(dataToAdd, value)

	}
	fmt.Println(dataToAdd)

}
