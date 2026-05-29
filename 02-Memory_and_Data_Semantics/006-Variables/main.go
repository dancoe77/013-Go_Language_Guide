package main

import "fmt"

func main() {
	var zz string = "########################################################"
	fmt.Println(zz)

	var aa string = "Memory is a byte - 00001010"
	var ab string = "Type of byte is integer which translates to 10"
	var ac string = "Type of byte is a color which translates to Red"
	var ad string = "Type is important because it dictates what the byte represents"
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(ad)
	fmt.Println(zz)

	//###############################################################################

	// Declare variables that are set to their zero value

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)
	fmt.Println(zz)

	// Declare variables and initialize.
	// Using the short variable declaration operator
	ba := 10
	bb := "hello"
	bc := 3.14159
	bd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", ba, ba)
	fmt.Printf("bb := 'hello' \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14.159 \t %T [%v]\n", bc, bc)
	fmt.Printf("dd := true \t %T [%v]\n", bd, bd)
	fmt.Println(zz)

	// Specify type and perform a conversion.
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
	fmt.Println(zz)
}

/*

 */
