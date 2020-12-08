// Package drsmodel provides Golang representations of DRS model definitions
//
// Module drsobject contains the DRSObject definition
package drsmodel

// DRSObject core object metadata class returned by the /objects/{object_id}
// controller. Contains metadata (description, checksums) as well as the
// access method to retrieve bytes
type DRSObject struct {
	AccessMethods []*AccessMethod   `json:"access_methods"`
	Aliases       []string          `json:"aliases"`
	Checksums     []*Checksum       `json:"checksums"`
	Contents      []*ContentsObject `json:"contents"`
	CreatedTime   string            `json:"created_time"`
	Description   string            `json:"description"`
	ID            string            `json:"id"`
	MimeType      string            `json:"mime_type"`
	Name          string            `json:"name"`
	SelfURI       string            `json:"self_uri"`
	Size          int64             `json:"size"`
	UpdatedTime   string            `json:"updated_time"`
	Version       string            `json:"version"`
}
