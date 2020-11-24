package randomize

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// PickSanta processes provided list to pair santas
func PickSanta(names []map[string]interface{}) string {
	log.SetPrefix("Randomizer-pickSanta: ")
	log.Output(1, "Picking Santas...")

	// Make sure we have an even amount of Santas to pair off
	if !(len(names)%2 == 0) {
		log.Fatal("There are an odd number of santas. Must be even, current santa count:", (len(names)))
	}

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names[rand.Intn(len(names))])
	fmt.Println("+++++++")

	// Loop by each map element
	for i, v := range names {
		personMatch := findMatch(i, len(names))
		fmt.Println("--------")
		fmt.Printf("|%s : %s|\n", v["name"], names[personMatch]["name"])
	}

	return "foo"
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
