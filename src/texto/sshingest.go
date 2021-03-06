package main

//ref: https://golangbot.com/read-files/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Read mode - line by line

	file, err := os.Open("sshinvalid.txt")
	if err != nil {
		fmt.Println("opening file error", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	fmt.Println("AS EACH LINE")
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
	var invalidUsersString []string

	//finding all lines with Invalid in line
	for _, eachline := range txtlines {
		fmt.Println(strings.Contains(eachline, "Invalid user"))
		if strings.Contains(eachline, "Invalid user") == true {
			invalidUsersString = append(invalidUsersString, eachline)

		}
	}

	var invalidUsers []string
	fmt.Println("UNVALID USER STRINGS FOUND:")
	for i := range invalidUsersString {
		fmt.Println(invalidUsersString[i])
		testArray := strings.Fields(invalidUsersString[i])

		invalidFound := false
		for _, v := range testArray {
			//fmt.Println(v)
			if v == "user" {
				fmt.Println("next iteration is user:")
				invalidFound = true
				continue
			}
			if invalidFound == true {
				fmt.Println("Found user name")
				fmt.Println(v)
				invalidUsers = append(invalidUsers, v)
				invalidFound = false
			}

		}
	}
	fmt.Println("======Invalid user list========")
	//fmt.Println(invalidUsers)
	for x := range invalidUsers {
		println("User: ", invalidUsers[x], " was found Invalid")
	}
}
