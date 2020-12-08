// Package drsmodel provides Golang representations of DRS model definitions
//
// Module checksum contains the checksum definition
package drsmodel

import (
	"github.com/ga4gh/drs-reference-server/internal/drsconstants"
)

// Checksum represents the hashed digest of DRS object bytes, including the
// raw checksum digest and the hashing algorithm used
type Checksum struct {
	Checksum string                    `json:"checksum"`
	Type     drsconstants.ChecksumType `json:"type"`
}

// NewChecksum instantiates a Checksum
func NewChecksum() *Checksum {
	return new(Checksum)
}

// SetChecksum sets raw checksum digest
func (checksum *Checksum) SetChecksum(sum string) {
	checksum.Checksum = sum
}

// GetChecksum returns the checksum digest
func (checksum *Checksum) GetChecksum() string {
	return checksum.Checksum
}

// SetType sets the 'type', or hashing algorithm used
func (checksum *Checksum) SetType(t drsconstants.ChecksumType) {
	checksum.Type = t
}

// GetType returns the 'type', or hashing algorithm used
func (checksum *Checksum) GetType() drsconstants.ChecksumType {
	return checksum.Type
}
