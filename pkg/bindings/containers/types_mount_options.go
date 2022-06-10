// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *MountOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *MountOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}
