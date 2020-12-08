// Package drsmodel provides Golang representations of DRS model definitions
//
// Module error contains the Error definition
package drsmodel

// Error objects are returned to the client when an HTTP error was encountered
// while processing the request (e.g. NotFound, BadRequest, etc.)
type Error struct {
	Message    string `json:"msg"`
	StatusCode int    `json:"status_code"`
}
