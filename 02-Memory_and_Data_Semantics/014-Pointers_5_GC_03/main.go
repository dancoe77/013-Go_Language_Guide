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

	an := "heap was 7 MB In-Use Before and 11 MB In-Use After"
	ao := "this is because there is application work happening in the other go routines"
	fmt.Println(an)
	fmt.Println(ao)
	fmt.Println(zz)

	ap := "When the GC misses a goal it will allocate a go routine to do mark assist"
	aq := "when this happens application throughput will be reduced from 75% to 50%"
	fmt.Println(ap)
	fmt.Println(aq)
	fmt.Println(zz)

	ar := "Application requires 10k requests which requires 5k GC maintaining a 40MB heap"
	as := "Determine what allocations the application is doing will allow someone to reduce"
	at := "the allocations/request that are happening, the goal being to reduce the total number of GCs"
	au := "to get the same amount of work done; ie. instead of 5k GCs make it 2500 GCs"
	fmt.Println(ar)
	fmt.Println(as)
	fmt.Println(at)
	fmt.Println(au)
	fmt.Println(zz)
}
