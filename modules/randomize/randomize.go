package randomize

import (
	"log"
	"math/rand"
	"time"
)

// MatchSanta processes provided list to pair santas
func MatchSanta(santaIndex int, names []map[string]interface{}, pickedNames []string) map[string]string {
	log.SetPrefix("Randomizer-MatchSanta: ")
	log.Printf("Finding match for %s...", names[santaIndex]["name"])
	// log.Printf("Already matched: %s", pickedNames)

	// Initialize and set match variable scope
	var validMatch bool
	var personMatch map[string]string

	// Validate that the match hasn't already been picked
	for validMatch == false {
		// Set match on each loop
		matchIndex := findMatchIndex(santaIndex, len(names))
		personMatch = ConvertMap(names[matchIndex])

		// Validate match criteria
		// Has not already been picked
		if hasBeenPicked(pickedNames, personMatch["name"]) == false {
			validMatch = true
		}
	}
	// log.Printf("Matched with: %s", personMatch["name"])

	return personMatch
}

func findMatchIndex(exclude int, total int) int {
	// Seed the randomizer
	rand.Seed(time.Now().UTC().UnixNano())

	// Loop until rand returns a number that is not excluded
	for {
		matchIndex := rand.Intn(total)
		if matchIndex != exclude {
			return matchIndex
		}
	}
}

// ConvertMap handles type assertion for map interface
func ConvertMap(originalMap interface{}) map[string]string {
	convertedMap := map[string]string{}
	for key, value := range originalMap.(map[string]interface{}) {
		convertedMap[key] = value.(string)
	}
	return convertedMap
}

// Find takes a slice and looks for an element in it. If found it will
// return true, otherwise it will return false.
func hasBeenPicked(slice []string, val interface{}) bool {
	// log.Printf("Checking if %s has already been picked...", val)
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
