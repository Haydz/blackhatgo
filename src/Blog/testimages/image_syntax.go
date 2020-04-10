// the idea is to write a script that will
/*
: create a text file with the syntax of including images in markdown format
so that a user can copy paste the syntact into their blog

steps:

1) identify the directory path for the images based on the root level of the project
example : /blog-folder/images/image-1.jpg becomes /images/image-1.jpg

2)take input for depth level, (1 folder down, 2 folder down)

3) take image file name append to directory path

4)must work on linux and windows
*/

/*
References
https://yourbasic.org/golang/list-files-in-directory/

*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// func doWindows() {
// 	directory, err := exec.Command("cmd", "/C", "echo %cd%").Output()
// 	if err != nil {
// 		fmt.Println("Error executing command", err)
// 	}
// 	return string(directory[:])
// }

// func doLinux() {
// 	directory, err := exec.Command("ls").Output()
// 	if err != nil {
// 		fmt.Println("Error executing command", err)
// 	}
// 	return string(directory[:])
// }

func main() {
	// var command string
	var osDirFormat string
	var directory []uint8
	if runtime.GOOS == "windows" {
		fmt.Println(runtime.GOOS)
		//command =
		osDirFormat = "\\"
		// err := nil
		directory, _ = exec.Command("cmd", "/C", "echo %cd%").Output()
		//fmt.Println(reflect.TypeOf(directory).String())
		//NEED TO FIX ERROR CHECKING
		// if err != nil {
		// 	fmt.Println("Error executing command", err)
		// }
	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
		directory, _ = exec.Command("pwd").Output()
		osDirFormat = "/"

	}
	// directory, err := exec.Command("cmd", "/C", "echo %cd%").Output()
	//identify current working directory
	// windows: echo %cd% linux :pwd:
	// directory, err := exec.Command("cmd", "/C", "echo %cd%").Output()
	// directory, err := exec.Command(command).Output()
	// if err != nil {
	// 	fmt.Println("Error executing command", err)
	// }
	//fmt.Println(string(directory[:]))
	//directory, err := exec.Command("cmd", "/C", "echo %cd%").Output()
	test := string(directory[:])
	//fmt.Println(reflect.TypeOf(test).String())
	fmt.Printf("%q\n", strings.Split(test, osDirFormat))

	//windows dirs split into slices
	testSplit := strings.Split(test, osDirFormat)
	//identifying current working Dir

	workingDir := testSplit[len(testSplit)-1]
	workingDir = strings.TrimSpace(workingDir) + osDirFormat
	fmt.Println("Working Directory is:", workingDir)
	workingDir = strings.TrimSpace(workingDir)
	//identify files in folder
	// fmt.Println("testing dir")
	// dir := path.Base(test)
	// fmt.Println(dir)

	//fmt.Println("On Unix:", filepath.SplitList(test))
	//fmt.Println(filepath.Base(test))

	filesList, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	//collecting current list of files
	var filesSlice []string
	for _, files := range filesList {
		// fmt.Println(files.Name())
		if files.Name() != "image_syntax.go" && files.Name() != "test.txt" && files.Name() != "image_syntax.exe" {
			// if files.Name() != ("image_syntax.go", "test.txt", "image_syntax.exe") {
			filesSlice = append(filesSlice, files.Name())
		}
	}

	//fmt.Println(filesSlice)
	outputFile, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer outputFile.Close()
	for i := range filesSlice {
		pathForFile := workingDir + filesSlice[i]
		//fmt.Printf("![%s](%s)\n", filesSlice[i], pathForFile)
		//stringToWrite := "![%s](%s)\n", filesSlice[i], pathForFile
		stringToWrite := "![" + filesSlice[i] + "](" + pathForFile + ")"
		fmt.Println(stringToWrite)
		_, err := io.WriteString(outputFile, stringToWrite+"\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	//return outputFile.Sync()

}

// 	//does not work for windows
// 	err2 := filepath.Walk(".",
// 		func(path string, info os.FileInfo, err2 error) error {
// 			if err2 != nil {
// 				return err2
// 			}
// 			fmt.Println(path)
// 			return nil
// 		})
// 	if err2 != nil {
// 		log.Println(err)
// 	}

// }
