// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how one needs to be careful when appending
// to a slice when you have a reference to an element.
package main

import "fmt"

type user struct {
	likes int
}

func main() {

	// Declare a slice of 3 users.
	users := make([]user, 3)

	// Share the user at index 1.
	shareUser := &users[1]

	// Add a like for the user that was shared.
	shareUser.likes++

	// Display the number of likes for all users.
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Add a new user.
	users = append(users, user{})

	// Add another like for the user that was shared.
	shareUser.likes++

	// Display the number of likes for all users.
	fmt.Println("********************************************")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Notice the last like has not been recorded. New backing array was implemented
	// length 4, capacity 6 but we never changed out the pointer so the likes for Ed
	// go back to the original backing array without being saved, and on reboot
	// anything held in memory disappears. Reference types use value semantics
	// for moving data around the programs but they use pointer semantics when
	// reading and writing.
}
