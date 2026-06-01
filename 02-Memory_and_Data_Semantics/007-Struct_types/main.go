package main

import (
	"fmt"
)

// example represents a type with different fields
type example struct {
	// Only use this if the situation calls for it
	// removing the extra padding bytes
	// counter int64
	// pi      float32
	// flag    bool
	// #######################################################
	flag bool
	// [3]byte padding
	// counter int32
	// [7]byte padding
	// counter int64
	counter int16
	pi      float32
	// [4]byte padding when using int64
}

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type nancy struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	zz := "#################################################################"
	fmt.Println(zz)

	// Declare a variable of type example set to its
	// zero value
	var e1 example

	// Display the value
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
	fmt.Println(zz)

	// Declare a variable of an anonymous type set
	// to its zero value
	var f1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", f1)

	// Declare a variable of an anonymous type and init
	// using a struct literal.
	f2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values.
	fmt.Printf("%+v\n", f2)
	fmt.Println("Flag", f2.flag)
	fmt.Println("Counter", f2.counter)
	fmt.Println("Pi", f2.pi)
	fmt.Println(zz)

	var b bill
	var n nancy
	// b = n
	// This will not work
	// major source of bugs
	// var b int
	// var b uint
	// this will work because f2 is not a named type
	// b = f2
	b = bill(n)
	fmt.Println(b, n)
	fmt.Println(zz)
}
