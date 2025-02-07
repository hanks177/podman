// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *RenameOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *RenameOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithName set field Name to given value
func (o *RenameOptions) WithName(value string) *RenameOptions {
	o.Name = &value
	return o
}

// GetName returns value of field Name
func (o *RenameOptions) GetName() string {
	if o.Name == nil {
		var z string
		return z
	}
	return *o.Name
}
