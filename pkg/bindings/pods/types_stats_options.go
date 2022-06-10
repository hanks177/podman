// Code generated by go generate; DO NOT EDIT.
package pods

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *StatsOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *StatsOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAll set field All to given value
func (o *StatsOptions) WithAll(value bool) *StatsOptions {
	o.All = &value
	return o
}

// GetAll returns value of field All
func (o *StatsOptions) GetAll() bool {
	if o.All == nil {
		var z bool
		return z
	}
	return *o.All
}
