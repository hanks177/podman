// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/hanks177/podman/v4/libpod/define"
	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *WaitOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *WaitOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithCondition set field Condition to given value
func (o *WaitOptions) WithCondition(value []define.ContainerStatus) *WaitOptions {
	o.Condition = value
	return o
}

// GetCondition returns value of field Condition
func (o *WaitOptions) GetCondition() []define.ContainerStatus {
	if o.Condition == nil {
		var z []define.ContainerStatus
		return z
	}
	return o.Condition
}

// WithInterval set field Interval to given value
func (o *WaitOptions) WithInterval(value string) *WaitOptions {
	o.Interval = &value
	return o
}

// GetInterval returns value of field Interval
func (o *WaitOptions) GetInterval() string {
	if o.Interval == nil {
		var z string
		return z
	}
	return *o.Interval
}
