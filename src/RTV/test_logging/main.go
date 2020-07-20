package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Results struct {
	CommandsList []string
	ID           int
	Command      string
	Output       string
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	output1 := &Results{
		ID:      1, // TODO: need to make this random for multiple children
		Output:  "command execution #1111",
		Command: "TEST 1",
	}

	output2 := &Results{
		ID:      2, // TODO: need to make this random for multiple children
		Output:  "command execution #222222",
		Command: "TEST 2",
	}
	fmt.Println("data to write to File)")
	fmt.Println(output1)
	fmt.Println(output2)

	fmt.Println("=====Testing Writing JSON=====")

	filename := "myFile.json"
	_ = checkFile(filename)

	// file, _ := json.MarshalIndent(output1, "", " ")
	file, _ := ioutil.ReadFile(filename)
	data := []Results{}
	json.Unmarshal(file, &data)
	data = append(data, *output1)
	dataBytes, _ := json.Marshal(&data)

	_ = ioutil.WriteFile(filename, dataBytes, 0644)
	// if err != nil {
	// 	logrus.Error(err)
	// }
	// dataBytes, _ := json.Marshal(output1)
	// _ = ioutil.WriteFile(filename, dataBytes, 0644)

	//////WORKS HERE//
	file, _ = ioutil.ReadFile(filename)
	// if err != nil {
	// 	logrus.Error(err)
	// }

	data = []Results{}

	json.Unmarshal(file, &data)

	data = append(data, *output2)

	dataBytes, _ = json.Marshal(data)
	// if err != nil {
	// 	logrus.Error(err)

	_ = ioutil.WriteFile(filename, dataBytes, 0644)

}
