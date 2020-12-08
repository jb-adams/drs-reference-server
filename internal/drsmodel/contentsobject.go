// Package drsmodel provides Golang representations of DRS model specifications
//
// Module contentsobject contains the ContentsObject definition
package drsmodel

// ContentsObject can contain either a single blob or a bundle of blobs,
// can be recursively unpackaged to access all DRS objects within a bundle
type ContentsObject struct {
	Contents []*ContentsObject `json:"contents"`
	DRSURI   []string          `json:"drs_uri"`
	ID       string            `json:"id"`
	Name     string            `json:"name"`
}
