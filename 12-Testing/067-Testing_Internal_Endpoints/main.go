// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that implements a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/dancoe77/013-Go_Language_Guide/12-Testing/067-Testing_Internal_Endpoints/handlers/handlers/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on: http://localhost:400")
	http.ListenAndServe(":4000", nil)
}
