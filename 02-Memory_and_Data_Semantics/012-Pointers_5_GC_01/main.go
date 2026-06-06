package main

import "fmt"

func main() {
	zz := "######################################################"
	fmt.Println(zz)

	aa := "Go's garbage collector is non-generational, non-compacting, concurrent, tri-color, mark and sweep"
	fmt.Println(aa)
	fmt.Println(zz)

	ab := "GOGC = 100"
	fmt.Println(ab)
	fmt.Println(zz)

	ac := "First GC = 4 MB"
	fmt.Println(ac)
	fmt.Println(zz)

	ad := "Stop_The_World < 100 microseconds"
	fmt.Println(ad)
	fmt.Println(zz)

	ae := "Phase 1; Mark - Start - Stop_The_World"
	af := "Phase 2; Marking values - Concurrent"
	ag := "Final Phase; Mark Termination - Stop_The_World"
	fmt.Println(ae)
	fmt.Println(af)
	fmt.Println(ag)
	fmt.Println(zz)

	ah := "100% CPU = 4 goroutines"
	fmt.Println(ah)
	fmt.Println(zz)
}
