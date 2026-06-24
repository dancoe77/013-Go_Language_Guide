// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to understand method sets.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email to  %s<%s>\n", u.name, u.email)
}

// duration is a named type with a base type of int.
type duration int

// notify implements the notifier interface
func (d *duration) notify() {
	fmt.Println("Sending  Notification in", *d)
}

func main() {

	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	// Values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	sendNotification(u)

	d := duration(42)
	d.notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n user) {
	n.notify()
}

/*
./example1.go:34: cannot use u (type user) as type notifier in argument to sendNotification:
user does not implement notifier (notify method has pointer receiver)

func sendNotification(n notifier) {
	n.notify()
}
*/

/*
duration(42).notify()

// ./example3.go:18: cannot call pointer method on duration(42)
	// ./example3.go:18: cannot take the address of duration(42)
*/
