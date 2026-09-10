// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// GOGC=off GOMAXPROCS=1 go test -bench . -benchtime 3s
// GOGC=off go test -bench . -benchtime 3s

// Tests to show the different performance based on concurrency with
// or without parallelism.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var numbers []int

func init() {
	rand.NewSource(time.Now().UnixNano())
	numbers = Generatelist(5e8)
	fmt.Printf("Processing %d numbers using %d goroutines on %d thread(s)\n", len(numbers), runtime.NumCPU(), runtime.GOMAXPROCS(0))
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Addconcurrent(runtime.NumCPU(), numbers)
	}
}

func BenchmarkSequentialAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(numbers)
	}
}

func BenchmarkConcurrentAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Addconcurrent(runtime.NumCPU(), numbers)
	}
}
