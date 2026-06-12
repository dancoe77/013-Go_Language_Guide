// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and iterate over
// arrays of different types.
package main

import "fmt"

func main() {

	// Declare an array of fixe strings that is initialized
	// to its zero value.
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// Iterate over the array of strings.
	// Current implementation uses value semantics, everyone gets a copy.
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// This implementation uses pointer semantics and indexes directly on the range
	/*
		for i := range fruits {
				fmt.Println(i, fruit)
			}
	*/

	// Declare an array of 4 integers that is initialized
	// with some values.
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers.
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}
