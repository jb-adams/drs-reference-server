// Package drsconstants provides constants throughout the program
//
// Module checksumtype provides allowable enumerated values for checksum types
package drsconstants

// ChecksumType enumerated values for checksum hashing algorithm names
type ChecksumType string

// ChecksumType enumeration
const (
	ChecksumTypeSHA256 ChecksumType = "sha-256"
	ChecksumTypeMD5    ChecksumType = "md5"
	ChecksumTypeEtag   ChecksumType = "etag"
	ChecksumTypeCrc32c ChecksumType = "crc32c"
)
