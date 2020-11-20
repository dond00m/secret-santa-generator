package randomize

import (
	"log"
)

// PickSanta processes provided list to pair santas
func PickSanta(names []map[string]interface{}) string {
	log.SetPrefix("Randomizer-pickSanta: ")
	log.Output(1, "Picking Santas...")

	if !(len(names)%2 == 0) {
		log.Fatal("There are an odd number of santas. Must be even, current santa count:", (len(names)))
	}
	// for _, v := range names {
	// 	fmt.Println("--------")
	// 	//Reading each value by its key
	// 	fmt.Printf("|Name: %s|\n", v["name"])
	// 	fmt.Printf("|Email: %s|\n", v["email"])
	// }

	return "foo"
}
