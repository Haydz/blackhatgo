package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//https://medium.com/@TobiasSchmidt89/effective-text-parsing-in-golang-163d13784288

func split(r rune) bool {
	return r == '[' || r == ']'
}

func main() {
	//need to open the file
	file, err := os.Open("sshinvalid.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
	}

	//we need to 'read' the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtLines []string
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		//fmt.Println(reflect.TypeOf(scanner.Text()))
		txtLines = append(txtLines, scanner.Text())

	}
	//https:medium.com/@TobiasSchmidt89/effective-text-parsing-in-golang-163d13784288

	// to break up with strings.Fields and strings.
	// try  string join

	//lets break up by space
	bySpace := strings.Fields(txtLines[0])
	fmt.Println(bySpace)

	//split string by 2 delimiters:
	twoDelims := strings.FieldsFunc(txtLines[1], split)

	for i := range twoDelims {
		fmt.Println(twoDelims[i])
	}

}
