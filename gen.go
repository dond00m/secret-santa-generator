package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("generate: ")

	log.Output(1, "JSON config")

	// Open JSON file
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	// Close file once function completes
	defer jsonFile.Close()

	// Capture JSON file contents
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Declare an empty interface of type array
	var results []map[string]interface{}

	// Unmarshal JSON to the interface
	json.Unmarshal([]byte(byteValue), &results)

	for _, v := range results {
		fmt.Println("--------")
		//Reading each value by its key
		fmt.Printf("|Name: %s|\n", v["name"])
		fmt.Printf("|Email: %s|\n", v["email"])
	}

}
