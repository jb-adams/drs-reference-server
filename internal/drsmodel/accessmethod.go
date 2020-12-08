// Package drsmodel provides Golang representations of DRS model definitions
//
// Module accessmethod contains the AccessMethod definition
package drsmodel

import (
	"github.com/ga4gh/drs-reference-server/internal/drsconstants"
)

// AccessMethod provides information on how to go about accessing object bytes
// from an external service, including the AccessURL
type AccessMethod struct {
	AccessID  string                        `json:"access_id"`
	AccessURL *AccessURL                    `json:"access_url"`
	Region    string                        `json:"region"`
	Type      drsconstants.AccessMethodType `json:"type"`
}

// NewAccessMethod instantiates an AccessMethod
func NewAccessMethod() *AccessMethod {
	return new(AccessMethod)
}

// SetAccessID sets access id
func (accessMethod *AccessMethod) SetAccessID(accessID string) {
	accessMethod.AccessID = accessID
}

// GetAccessID returns access id
func (accessMethod *AccessMethod) GetAccessID() string {
	return accessMethod.AccessID
}

// SetAccessURL sets access URL
func (accessMethod *AccessMethod) SetAccessURL(accessURL *AccessURL) {
	accessMethod.AccessURL = accessURL
}

// GetAccessURL returns access URL
func (accessMethod *AccessMethod) GetAccessURL() *AccessURL {
	return accessMethod.AccessURL
}

// SetRegion sets region
func (accessMethod *AccessMethod) SetRegion(region string) {
	accessMethod.Region = region
}

// GetRegion returns region
func (accessMethod *AccessMethod) GetRegion() string {
	return accessMethod.Region
}

// SetType sets access method type
func (accessMethod *AccessMethod) SetType(t drsconstants.AccessMethodType) {
	accessMethod.Type = t
}

// GetType returns access method type
func (accessMethod *AccessMethod) GetType() drsconstants.AccessMethodType {
	return accessMethod.Type
}
