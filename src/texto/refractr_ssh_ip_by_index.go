package main
// mext idea to do stats, like top IP address trying logins and what users
//ref: https://golangbot.com/read-files/
// https://bash-prompt.net/guides/using-logs-1/
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//IDEA: maybe make the searching an individual function?
func main() {

	file, err := os.Open("sshinvalid.txt") // opening file
	if err != nil {                        // error catching
		fmt.Println("opening file error", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ipMapping = map[string]string{}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Invalid user") == true { // we want Invalid User line
			// I want to find identify the IP addresses and Usernames
			testArray := strings.Fields(scanner.Text())
			userName := testArray[7]
			if len(testArray) == 10 { //catches line that is missing IP address at end.
				ipFound := testArray[9]
				ipMapping[userName] = ipFound
			}

		} else if strings.Contains(scanner.Text(), "logname=") == true {
			re := regexp.MustCompile("user=\\w+")
			if re.FindString(scanner.Text()) != "" {
				fmt.Println(re.FindString(scanner.Text()))
				fmt.Println(strings.Replace(re.FindString(scanner.Text()), "user=", "", -1))
			}
		}

	}
	file.Close()
	fmt.Println("======Invalid User and Origin IP list========")
	//fmt.Println(invalidUsers)
	for key, value := range ipMapping {
		println("User: ", key, " Originated from:", value)
	}

}
