// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the default error type is implemented.
package main

import "fmt"

// http://golang.org/src/pkg/builtin/#error
type error interface {
	Error() string
}

// http://golang,org/src/pkg/errors/errors.go
type errorString struct {
	s string
}

// http://golang,org/src/pkg/errors/errors.go
func (e *errorString) Error() string {
	return e.s
}

// http://golang,org/src/pkg/errors/errors.go
// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

func main() {
	zz := "############################################################"
	fmt.Println(zz)

	aa := "Systems cannot be developed assuming that human beings will be able to write millions without making mistakes,"
	ab := "and debugging alone is not an efficient way to develop reliable systems."
	ac := "Al Aho (inventor of AWK)"
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(zz)

	if err := webCall(); err != nil {
		fmt.Println(err)
		fmt.Println(zz)
		return
	}

	fmt.Println("Life is good")
	fmt.Println(zz)

}

// webCall performs a web operation.
func webCall() error {
	return New("Bad Request")
}
