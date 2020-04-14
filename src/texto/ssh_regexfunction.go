package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func regexParser(regex string, line string) string {
	re := regexp.MustCompile(regex)

	re.FindString(line)

	return re.FindString(line)

}

func main() {

	file, err := os.Open("sshinvalid.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
	}

	//need to read in file
	scanner := bufio.NewScanner(file)
	//want to split by line
	scanner.Split(bufio.ScanLines)

	fmt.Println(scanner.Scan())

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		word := regexParser("user=\\w+", scanner.Text())
		if word != "" {
			fmt.Println(word)
		}
	}

}
