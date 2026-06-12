// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the capacity of the slice
// is not available for use.
package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements.
	// fruits := make([]string, 5)
	// Create a slice with a length of 5 elements and a capacity of 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// You can't access an index of a slice beyond its length.
	// fruits[5] = "Runtime error"

	// Error: panic: runtime error: index out of range
	// fmt.Println(fruits)

	inspectSlice(fruits)
}

// inspectSlice exposes the slice header for review
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}
