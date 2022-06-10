// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *StopOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *StopOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithIgnore set field Ignore to given value
func (o *StopOptions) WithIgnore(value bool) *StopOptions {
	o.Ignore = &value
	return o
}

// GetIgnore returns value of field Ignore
func (o *StopOptions) GetIgnore() bool {
	if o.Ignore == nil {
		var z bool
		return z
	}
	return *o.Ignore
}

// WithTimeout set field Timeout to given value
func (o *StopOptions) WithTimeout(value uint) *StopOptions {
	o.Timeout = &value
	return o
}

// GetTimeout returns value of field Timeout
func (o *StopOptions) GetTimeout() uint {
	if o.Timeout == nil {
		var z uint
		return z
	}
	return *o.Timeout
}
