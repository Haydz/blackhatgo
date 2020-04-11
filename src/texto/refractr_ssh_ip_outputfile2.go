package main

//ref: https://golangbot.com/read-files/
// https://bash-prompt.net/guides/using-logs-1/
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

//counting valid users function
func countValidUsers(users []string) int {

	return len(users)
}

//counting key and valu numbers
func countMap(test map[string]string) (int, int) {
	keyNum := 0
	valueNum := 0

	for _, value := range test {
		if value == "NA" {
			// fmt.Println("found user WIHOUT IP, valuenum: ", valueNum)
			keyNum = keyNum + 1

		} else {
			valueNum = valueNum + 1
			// fmt.Println("found user with IP, valuenum: ", valueNum)
			keyNum = keyNum + 1
		}
	}
	return keyNum, valueNum
}

//IDEA: maybe make the searching an individual function?
func main() {

	file, err := os.Open("sshinvalid.txt") // opening file
	if err != nil {                        // error catching
		fmt.Println("opening file error", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ipMapping = map[string]string{}
	var validIpUsers = map[string]string{}
	var validUsers []string

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "Invalid user") == true { // we want Invalid User line
			// I want to find identify the IP addresses and Usernames
			testArray := strings.Fields(scanner.Text())
			userName := testArray[7]
			if len(testArray) == 10 { //catches line that is missing IP address at end.
				ipFound := testArray[9]
				ipMapping[userName] = ipFound
			}

		} else if strings.Contains(scanner.Text(), "logname=") == true { //finding users that were legitimate
			re := regexp.MustCompile("user=\\w+")
			ipAddress := regexp.MustCompile("\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b")
			fmt.Println(ipAddress.FindString(scanner.Text()))
			if re.FindString(scanner.Text()) != "" {
				//fmt.Println(re.FindString(scanner.Text()))
				singleUser := strings.Replace(re.FindString(scanner.Text()), "user=", "", -1)
				compare := (contains(validUsers, singleUser))
				fmt.Println(compare)
				if compare != true {
					validUsers = append(validUsers, singleUser)
					if ipAddress.FindString(scanner.Text()) == "" {
						validIpUsers[singleUser] = "NA"
					} else {
						validIpUsers[singleUser] = ipAddress.FindString(scanner.Text())
					}
					//}
				}
			}
		}
	}

	file.Close()
	fmt.Println("======Invalid User and Origin IP list========")
	//fmt.Println(invalidUsers)
	for key, value := range ipMapping {
		println("User: ", key, " Originated from:", value)
	}
	fmt.Println("Valid users slice: ", validUsers)
	fmt.Println("MAP of Valid Users and IP addresses", validIpUsers)
	fmt.Println("The number of valid users from validUsers Array: ", countValidUsers(validUsers))
	mapUsers, mapIPs := countMap(validIpUsers)
	fmt.Println("the number of Valid users according to the MAP:", mapUsers)
	fmt.Println("number of IP addresses from the MAP, minus NA:", mapIPs)
}
