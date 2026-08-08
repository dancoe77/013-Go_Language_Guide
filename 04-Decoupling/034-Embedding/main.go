// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how what we are doing is NOT embedding
// a type but just using a type as a field.
package main

import "fmt"

// notifier is an interface that defined notification
// type behavior
type notifier interface {
	notify()
}

// user defines a user in the program
type user struct {
	name  string
	email string
}

// notify implements a method notifies users
// of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileges.
type admin struct {
	//person user // NOT Embedding
	user  // Embedded Type
	level string
}

// notify implements a method notifies admins
// of different events.
func (a *admin) notify() {
	fmt.Printf("Sending admin Email To %s<%s>\n", a.name, a.email)
}

func main() {
	zz := "###########################################################"
	fmt.Println(zz)

	// Create an admin user.
	ad := admin{
		user: user{ //person: user
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	// We can access fields methods.
	//ad.person.notify()
	// We can access the inner type's method directly.
	ad.user.notify()
	fmt.Println(zz)
	// The inner type's method is promoted.
	ad.notify()
	fmt.Println(zz)

	// Send the admin user a notification .
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification(&ad)
	fmt.Println(zz)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
