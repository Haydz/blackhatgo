package main

//ref: https://golangbot.com/read-files/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

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

	var invalidUsers []string
	//finding all lines with Invalid in line
	for _, eachline := range txtlines {
		//fmt.Println(strings.Contains(eachline, "Invalid user"))
		if strings.Contains(eachline, "Invalid user") == true {
			//invalidUsersString = append(invalidUsersString, eachline)
			testArray := strings.Fields(eachline)
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
	}
	fmt.Println("======Invalid user list========")
	//fmt.Println(invalidUsers)
	for x := range invalidUsers {
		println("User: ", invalidUsers[x], " was found Invalid")
	}

	// I want to find identify the IP addresses and Usernames

}
