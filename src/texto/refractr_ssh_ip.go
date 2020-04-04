package main

//ref: https://golangbot.com/read-files/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
//IDEA: maybe make the searching an individual function?
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
	var ipMapping = map[string]string{}

	var invalidUsers []string
	//finding all lines with Invalid in line
	for _, eachline := range txtlines {

		if strings.Contains(eachline, "Invalid user") == true {
			// I want to find identify the IP addresses and Usernames
			testArray := strings.Fields(eachline)
			invalidFound := false
			ipFound := false
			userName := ""
			for _, v := range testArray {

				//fmt.Println(v)

				if v == "user" {
					//fmt.Println("next iteration is user:")
					invalidFound = true
					continue
				}
				if invalidFound == true {
					userName = v
					//fmt.Println("Found user name")
					//fmt.Println(v)
					invalidUsers = append(invalidUsers, v)

					invalidFound = false
					// need to loop to find IP addess
					for _, v = range testArray {
						if v == "from" {
							//next iteration is IP address
							ipFound = true
							continue
						}
						if ipFound == true {
							fmt.Println("username:", userName, " has IP of ", v)
							ipFound = false
							// add to MAP
							ipMapping[userName] = v

						}
					}
				}
			}
		}
	}
	// fmt.Println("======Invalid user list========")
	// //fmt.Println(invalidUsers)
	// for x := range invalidUsers {
	// 	println("User: ", invalidUsers[x], " was found Invalid")
	// }
	fmt.Println("======Invalid User and Origin IP list========")
	//fmt.Println(invalidUsers)
	for key, value := range ipMapping {
		println("User: ", key, " Originated from:", value)

	}
}
