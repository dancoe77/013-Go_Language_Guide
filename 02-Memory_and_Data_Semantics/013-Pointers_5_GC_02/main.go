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

	ai := "Go 1.14 or > has a preemptive scheduler built into the runtime"
	fmt.Println(ai)
	fmt.Println(zz)

	aj := "Garbage Collector will use up to 25% of the available CPU to perform work"
	ak := "75% CPU = 3 goroutines"
	fmt.Println(aj)
	fmt.Println(ak)
	fmt.Println(zz)

	al := "Garbage collector is tricolor, mark and sweep"
	fmt.Println(al)
	fmt.Println(zz)

	am := "GODEBUG=gctrace=1 ./app"
	fmt.Println(am)
	fmt.Println(zz)
}
