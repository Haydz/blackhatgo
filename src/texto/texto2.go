package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//need to open the file
	file, err := os.Open("sshinvalid.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
	}

	//we need to 'read' the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//fmt.Println(reflect.TypeOf(scanner.Text()))

	}

}
