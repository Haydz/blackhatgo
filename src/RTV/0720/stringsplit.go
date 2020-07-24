package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	test := "127.0.0.1:1111"

	// test2 := strings.Split(test, ":")

	patternIP := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$"

	res2, e := regexp.MatchString(patternIP, test)
	fmt.Println("result:", res2, e, test)

	test2 := strings.Split(test, ":")

	portCheck, _ := strconv.Atoi(test2[1])
	fmt.Println(portCheck < 65535)
}
