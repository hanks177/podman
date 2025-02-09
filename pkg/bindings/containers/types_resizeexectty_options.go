// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *ResizeExecTTYOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *ResizeExecTTYOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithHeight set field Height to given value
func (o *ResizeExecTTYOptions) WithHeight(value int) *ResizeExecTTYOptions {
	o.Height = &value
	return o
}

// GetHeight returns value of field Height
func (o *ResizeExecTTYOptions) GetHeight() int {
	if o.Height == nil {
		var z int
		return z
	}
	return *o.Height
}

// WithWidth set field Width to given value
func (o *ResizeExecTTYOptions) WithWidth(value int) *ResizeExecTTYOptions {
	o.Width = &value
	return o
}

// GetWidth returns value of field Width
func (o *ResizeExecTTYOptions) GetWidth() int {
	if o.Width == nil {
		var z int
		return z
	}
	return *o.Width
}
