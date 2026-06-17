// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to initialize a map, write to
// it, then read and delete from it.
package main

import (
	"fmt"
	"sort"
)

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

type player struct {
	name  string
	score int
}

// type users []user

func main() {

	zz := "##########################################################################"

	// Declare and make a map that stores values
	// of type user with a key of type string.
	users := make(map[string]user)

	// Add key/value pairs to the map.
	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Michael", "Jackson"}

	// Read the value at a specific key.
	mouse := users["Mouse"]
	fmt.Printf("%+v\n", mouse)

	// Replace the value at the Mouse key.
	users["Mouse"] = user{"Jerry", "Mouse"}

	// Read the Mouse key again.
	fmt.Printf("%+v\n", mouse)

	// Delete the value at a specific key.
	delete(users, "Roy")

	// Check the length of the map. There are only 3 elements.
	fmt.Println(len(users))

	// It is safe to delete an absent key.
	delete(users, "Roy")

	fmt.Println("Goodbye.")

	fmt.Println(zz)

	// Create a map to track scores for players in a game.
	scores := make(map[string]int)

	// Read the element at key "anna". It is absent so we get
	// the zero-value for this map's value type.
	score := scores["anna"]
	fmt.Println("Score:", score)

	// If we need to check for the presence of a key we use
	// a 2 variable assignment. The 2nd variable is a bool.
	score, ok := scores["anna"]
	fmt.Println("Score:", score, "Present:", ok)

	// We can leverage the zero-value behavior to write
	// convenient code like this:
	scores["anna"]++

	// Without this behavior we would have to code in a
	// defensive way like this:
	if n, ok := scores["anna"]; ok {
		scores["anna"] = n + 1
	} else {
		scores["anna"] = 1
	}

	score, ok = scores["anna"]
	fmt.Println("Score:", score, "Present:", ok)

	fmt.Println(zz)

	// u := make(map[users]int)

	// ./example3.go:22: invalid map key type users

	// Iterate over the map.
	//for key, value := range u {
	//	fmt.Println(key, value)
	//}

	// Declare and initialize the map with values.
	users1 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Iterate over the map printing each key and value.
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterate over the map printing just the keys.
	// Notice the results are different.
	for key := range users1 {
		fmt.Println(key)
	}
	fmt.Println(zz)

	// Pull the keys from the map.
	var keys []string
	for key := range users1 {
		keys = append(keys, key)
	}
	// Sort the keys alphabetically
	sort.Strings(keys)

	// Walt through the keys and pull each value from the map.
	for _, key := range keys {
		fmt.Println(key, users1[key])
	}
	fmt.Println(zz)

	// Declare a map with initial values using a map literal.
	players := map[string]player{
		"anna":  {"Anna", 42},
		"jacob": {"Jacob", 21},
	}

	// Iterate over the map printing each key and value.
	for key, value := range players {
		fmt.Println(key, value)
	}

	// Trying to take the address of a map element fails.
	// anna := &players["anna"]
	// anna.score++

	// ./example4.go:23:10: cannot take the address of players["anna"]

	// Instead take the element, modify it, and put it back.
	player := players["anna"]
	player.score++
	players["anna"] = player

	// Iterate over the map printing each key and value.
	for key, value := range players {
		fmt.Println(key, value)
	}
	fmt.Println(zz)

	// Initialize a map with values.
	scores1 := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Pass the map to a fuction to perform some mutation.
	double(scores1, "anna")

	// See the change is visible in our map.
	fmt.Println("Score:", scores1["anna"])

	fmt.Println(zz)
}

// Double finds the score for a specific player and
// multiples it by 2.
func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}
