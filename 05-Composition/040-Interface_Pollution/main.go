// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This is an example that creates interface pollution
// by improperly using an interface when one is not needed.
package main

/*
// Server defines a contract for tcp servers.
type Server interface {
	Start() error
	Stop() error
	Wait() error
}
*/

// Server is our Server implementation.
type Server struct {
	host string

	// Pretend there are more fields.
}

// NewServer returns an interface value of type Server
// with a server implementation.
func NewServer(host string) *Server {

	// SMELL - Storing an unexported type pointer in the interface.
	return &Server{host}
}

// Start allows the server to begin to accept requests.
func (s *Server) Start() error {

	// Pretend there is a specific implementation
	return nil
}

// Stop shuts the server down.
func (s *Server) Stop() error {

	// Pretend there is a specific implementation.
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *Server) Wait() error {

	// Pretend there is a specific implementation.
	return nil
}

func main() {
	// Create a new Server.
	srv := NewServer("localhost")

	// Use the API.
	srv.Start()
	srv.Stop()
	srv.Wait()
}

// #################################################################################

// NOTES:

// Smells:
//  * The package declares an interface that matches the entire API of its own concrete type.
//  * The interface is exported but the concrete type is unexported.
//  * The factory function returns the interface value with the unexported concrete type value inside.
//  * The interface can be removed and nothing changes for the user of the API.
//  * The interface is not decoupling the API from change.

// ##################################################################################

// NOTES:

// Here are some guidelines around interface pollution:
// * Use an interface:
//      * When users of the API need to provide an implementation detail.
//      * When API’s have multiple implementations that need to be maintained.
//      * When parts of the API that can change have been identified and require decoupling.
// * Question an interface:
//      * When its only purpose is for writing testable API’s (write usable API’s first).
//      * When it’s not providing support for the API to decouple from change.
//      * When it's not clear how the interface makes the code better.
