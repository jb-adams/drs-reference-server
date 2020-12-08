// Package drsmodel provides Golang representations of DRS model definitions
//
// Module accessurl contains the AccessURL definition
package drsmodel

// AccessURL provides information on how to access object bytes from an
// extenal service, local file, cloud bucket, etc. Contains the fully resolvable
// URL and any headers that must be passed in the request
type AccessURL struct {
	Headers []string `json:"headers"`
	URL     string   `json:"url"`
}

// NewAccessURL instantiates an AccessURL
func NewAccessURL() *AccessURL {
	return new(AccessURL)
}

// SetHeaders sets headers to be passed in HTTP request
func (accessURL *AccessURL) SetHeaders(headers []string) {
	accessURL.Headers = headers
}

// AddHeader appends a single header to the list of headers
func (accessURL *AccessURL) AddHeader(header string) {
	accessURL.Headers = append(accessURL.Headers, header)
}

// GetHeaders returns all headers
func (accessURL *AccessURL) GetHeaders() []string {
	return accessURL.Headers
}

// SetURL sets the URL to object bytes
func (accessURL *AccessURL) SetURL(url string) {
	accessURL.URL = url
}

// GetURL returns the URL to object bytes
func (accessURL *AccessURL) GetURL() string {
	return accessURL.URL
}
