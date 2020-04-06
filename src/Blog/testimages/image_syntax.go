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
	"reflect"
	"strings"
)

func main() {

	/*
			    if runtime.GOOS == "windows" {
		        fmt.Println("Can't Execute this on a windows machine")
		    } else {
		        execute()
			}
	*/

	//identify current working directory
	// windows: echo %cd% linux :pwd:

	directory, err := exec.Command("cmd", "/C", "echo %cd%").Output()
	if err != nil {
		fmt.Println("Error executing command", err)
	}
	fmt.Println(string(directory[:]))
	test := string(directory[:])
	fmt.Println(reflect.TypeOf(test).String())
	fmt.Printf("%q\n", strings.Split(test, "\\"))

	//windows dirs split into slices
	testSplit := strings.Split(test, "\\")
	//identifying current working Dir

	winStruct := "\\"
	// linStruct := "/"
	workingDir := testSplit[len(testSplit)-1]
	workingDir = strings.TrimSpace(workingDir) + winStruct
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
		if files.Name() != "image_syntax.go" && files.Name() != "test.txt" {
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
