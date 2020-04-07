package main

import (
	"fmt"
	"runtime"
)

func main() {

	// test if windows or linux
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		fmt.Println("Windows")
	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}

}
