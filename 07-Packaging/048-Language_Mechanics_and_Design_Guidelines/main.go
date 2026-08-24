package main

import "fmt"

func main() {
	zz := "##################################################################################################"
	fmt.Println(zz)

	aa := "I think the real problem with C is that it doesn't give you enough mechanisms for structuring really big programs,"
	ab := "for creating 'firewalls' within programs so you can keep the various pieces apart."
	ac := "It's not that you can't do all  of these things, that you can't simulate object-oriented programming or other methodology you want in C."
	ad := "You can simulate it, but the compiler, the language itself isn't giving you any help."
	ae := "- Brian Kernighan"
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(ad)
	fmt.Println(ae)
	fmt.Println(zz)

	ba := "Packaging directly conflicts with how we have been taught to organize source code in other languages."
	bb := "In other languages, packaging is a feature that you can choose to use or ignore."
	bc := "You can think of packaging as applying the idea fo microservices on a source tree."
	bd := "All packages are 'first class,'and the only hierarchy is what you define in the source tree for your project."
	be := "There needs to be a way to 'open' parts of the package to the outside world."
	bf := "Two packages can't cross-import each other. Imports are a one way street."
	fmt.Println(ba)
	fmt.Println(bb)
	fmt.Println(bc)
	fmt.Println(bd)
	fmt.Println(be)
	fmt.Println(bf)
	fmt.Println(zz)

	ca := "To be purposeful, packages must provide, not contain."
	cb := "- Packages must be named with the intent to describe what it provides."
	cc := "- Packages must not become a dumping ground of disparate concerns."
	fmt.Println(ca)
	fmt.Printf("\t%v\n", cb)
	fmt.Printf("\t%v\n", cc)
	fmt.Println(zz)

	da := "To be usable, packages must be designed with the user as their focus."
	db := "- Packages must be intuitive and simple to use."
	dc := "- Packages must respect their impact on resources and performance."
	dd := "- Packages must protect the user's application from cascading changes."
	de := "- Packages must prevent the need for type assertions to the concrete."
	df := "- Packages must reduce, minimize and simplify its code base."
	fmt.Println(da)
	fmt.Printf("\t%v\n", db)
	fmt.Printf("\t%v\n", dc)
	fmt.Printf("\t%v\n", dd)
	fmt.Printf("\t%v\n", de)
	fmt.Printf("\t%v\n", df)
	fmt.Println(zz)

	ea := "To be portable, packages must be designed with reusability in mind."
	eb := "- Packages must aspire for the highest level of portability."
	ec := "- Packages must reduce setting policy when it's reasonable and practical."
	ed := "- Packages must not become a single point of dependency."
	fmt.Println(ea)
	fmt.Printf("\t%v\n", eb)
	fmt.Printf("\t%v\n", ec)
	fmt.Printf("\t%v\n", ed)
	fmt.Println(zz)
}
