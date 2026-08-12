// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This is an example of using type hierarchies with a OOP pattern.
// This is not something we want to do in Go. Go does not have the
// concept of sub-typing. All types are their own and the concepts of
// base and derived types do not exist in Go. This pattern does not
// provide a good design principle in a Go program.
package main

import "fmt"

/*
// Animal contains all the base fields for animals.
type Animal struct {
	Name     string
	IsMammal bool
}

// Speak provides generic behavior for all animals and
// how they speak.
func (a *Animal) Speak() {
	fmt.Printf("UGH! My name is %s, it is %t I am a mammal\n", a.Name, a.IsMammal)
}
*/

// Speaker provides a common behavior for all concrete types
// to follow if they want to be a part of this group. This
// is a contract for these contract types to follow.
type Speaker interface {
	Speak()
}

// Dog contains everything a Dog needs.
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d *Dog) Speak() {
	fmt.Printf("Woof! My name is %s, it is %t I am a mammal with a pack factor of %d.\n", d.Name, d.IsMammal, d.PackFactor)
}

// Cat contains everything a Cat needs.
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (c *Cat) Speak() {
	fmt.Printf("Meow! My name is %s, it is %t I am a mammal with a climb factor of %d.\n", c.Name, c.IsMammal, c.ClimbFactor)
}

func main() {
	zz := "#####################################################"
	fmt.Println(zz)

	aa := "A good API is not just easy to use but also hard to misuse"
	ab := "- JBD"
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(zz)

	ba := "You can always embed, but you cannot decompose big interfaces once they are out there."
	bb := "Keep interfaces small."
	bc := "- JBD"
	fmt.Println(ba)
	fmt.Println(bb)
	fmt.Println(bc)
	fmt.Println(zz)

	ca := "Don't design with interfaces, discover them."
	cb := " -Rob Pike"
	fmt.Println(ca)
	fmt.Println(cb)
	fmt.Println(zz)

	da := "Duplication is far cheaper than the wrong abstraction."
	db := "- Sandi Metz"
	fmt.Println(da)
	fmt.Println(db)
	fmt.Println(zz)

	/*
		// Create a list of Animals that know how to speak.
		animals := []Animal{

			// Create a Dog by initializing its Animal parts
			// and then it's specific Dog attributes.
			Dog{
				Animal: Animal{
					Name:     "Fido",
					IsMammal: true,
				},
				PackFactor: 5,
			},

			// Create a  Cat by initializing its Animal parts
			// and then its specific Cat attributes.
			Cat{
				Animal: Animal{
					Name:     "Milo",
					IsMammal: true,
				},
				ClimbFactor: 4,
			},
		}

		// Have the Animals speak.
		for _, animal := range animals {
			animal.Speak()
		}
	*/

	// Create a list of Animals that know how to speak.
	speakers := []Speaker{

		// Create a Dog by initializing its Animal parts
		// and then its specific Dog attributes.
		&Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts
		// and then its specific Cat attributes.
		&Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, spkr := range speakers {
		spkr.Speak()
	}
	fmt.Println(zz)
}

// =============================================================================

// NOTES:

// Smells:
// 	* The Animal type is providing an abstraction layer of reusable state.
// 	* The program never needs to create or solely use a value of type Animal.
// 	* The implementation of the Speak method for the Animal type is a generalization.
// 	* The Speak method for the Animal type is never going to be called.

// =============================================================================

// NOTES:

// Here are some guidelines around declaring types:
// 	* Declare types that represent something new or unique.
// 	* Validate that a value of any type is created or used on its own.
// 	* Embed types to reuse existing behaviors you need to satisfy.
// 	* Question types that are an alias or abstraction for an existing type.
// 	* Question types whose sole purpose is to share common state.
