// Sample program to teach the mechanics of escape analysis.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application
func main() {
	zz := "##########################################"
	fmt.Println(zz)

	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)

	fmt.Println(zz)
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
//go:noinline

func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)
	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
//
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	return &u
}

/*
// See escape analysis and inlining decisions.

$ go build -gcflags -m=2
# github.com/ardanlabs/gotraining/topics/go/language/pointers/example4
./example4.go:24:6: cannot inline createUserV1: marked go:noinline
./example4.go:38:6: cannot inline createUserV2: marked go:noinline
./example4.go:14:6: cannot inline main: non-leaf function
./example4.go:30:16: createUserV1 &u does not escape
./example4.go:46:9: &u escapes to heap
./example4.go:46:9:      from ~r0 (return) at ./example4.go:46:2
./example4.go:46:9: moved to heap: u
./example4.go:44:16: createUserV2 &u does not escape
./example4.go:18:16: main &u1 does not escape
./example4.go:18:27: main &u2 does not escape

*/
