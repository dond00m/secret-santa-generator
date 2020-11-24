package randomize

import (
	"log"
	"math/rand"
	"time"
)

// MatchSanta processes provided list to pair santas
func MatchSanta(santaIndex int, names []map[string]interface{}) int {
	log.SetPrefix("Randomizer-pickSanta: ")
	log.Printf("Finding match for %s...", names[santaIndex]["name"])

	matchIndex := findMatch(santaIndex, len(names))
	return matchIndex
}

func findMatch(exclude int, total int) int {
	// Seed the randomizer
	rand.Seed(time.Now().UTC().UnixNano())

	// Loop until rand returns a number that is not excluded
	for {
		selection := rand.Intn(total)
		if selection != exclude {
			return selection
		}
	}

}
