// Package drsconstants provides constants through the program
//
// Module accessmethodtype provides allowable enum values for URL schemes
// in access URLs
package drsconstants

// AccessMethodType enumerated values for URL schemes in access URLs
type AccessMethodType string

// AccessMethodType enumeration
const (
	AccessMethodTypeHTTPS AccessMethodType = "https"
	AccessMethodTypeFile  AccessMethodType = "file"
)
