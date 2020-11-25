package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/notify"
	"local/randomize"
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
		// Initialize pair struct
		p := notify.SantaPair{}
		// Type assertion record to map
		santa := randomize.ConvertMap(v)
		// Get match record
		match := randomize.MatchSanta(i, santas, matchedSantas)
		// Update matched santa slice
		matchedSantas = append(matchedSantas, match["name"])
		log.Printf("%s is the secret santa for: %s", santa["name"], match["name"])
		// Combine santa and their match to struct
		p.SantaName = santa["name"]
		p.SantaEmail = santa["email"]
		p.ReceipentName = match["name"]
		p.ReceipentEmail = match["email"]
		p.ReceipentAddress = match["address"]
		p.ReceipentWishlist = match["wishlist"]
		notify.SendEmail(p)
	}
}
