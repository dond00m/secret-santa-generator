package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"local/randomize"
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

	// Declare an empty map interface
	var santas []map[string]interface{}

	// Unmarshal JSON to the interface
	json.Unmarshal([]byte(byteValue), &santas)

	// Make sure we have an even amount of Santas to pair off
	if !(len(santas)%2 == 0) {
		log.Fatal("There are an odd number of santas. Must be even, current santa count:", (len(santas)))
	}
	// Initialize slice to track who has already been matched
	matchedSantas := make([]string, 0)

	for i, v := range santas {
		fmt.Println("--------")
		fmt.Println(matchedSantas)
		personMatch := convertMap(santas[randomize.MatchSanta(i, santas)])
		if Find(matchedSantas, personMatch["name"]) == false {
			log.Output(1, "Valid match!")
			fmt.Printf("|%s : %v,%s|\n", v["name"], personMatch["name"], personMatch["email"])
			matchedSantas = append(matchedSantas, personMatch["name"])
		} else {
			log.Printf("%s already matched, re-rolling", personMatch["name"])
			personMatch := convertMap(santas[randomize.MatchSanta(i, santas)])
			fmt.Printf("|%s : %v,%s|\n", v["name"], personMatch["name"], personMatch["email"])

		}
	}
}

// Handle type assertion for map interface
func convertMap(originalMap interface{}) map[string]string {
	convertedMap := map[string]string{}
	for key, value := range originalMap.(map[string]interface{}) {
		convertedMap[key] = value.(string)
	}
	return convertedMap
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val interface{}) bool {
	log.Printf("Checking if %s has already been picked..", val)
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
