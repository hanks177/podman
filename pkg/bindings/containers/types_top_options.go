// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *TopOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *TopOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithDescriptors set field Descriptors to given value
func (o *TopOptions) WithDescriptors(value []string) *TopOptions {
	o.Descriptors = &value
	return o
}

// GetDescriptors returns value of field Descriptors
func (o *TopOptions) GetDescriptors() []string {
	if o.Descriptors == nil {
		var z []string
		return z
	}
	return *o.Descriptors
}
