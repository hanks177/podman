// Code generated by go generate; DO NOT EDIT.
package images

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *SearchOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *SearchOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAuthfile set field Authfile to given value
func (o *SearchOptions) WithAuthfile(value string) *SearchOptions {
	o.Authfile = &value
	return o
}

// GetAuthfile returns value of field Authfile
func (o *SearchOptions) GetAuthfile() string {
	if o.Authfile == nil {
		var z string
		return z
	}
	return *o.Authfile
}

// WithFilters set field Filters to given value
func (o *SearchOptions) WithFilters(value map[string][]string) *SearchOptions {
	o.Filters = value
	return o
}

// GetFilters returns value of field Filters
func (o *SearchOptions) GetFilters() map[string][]string {
	if o.Filters == nil {
		var z map[string][]string
		return z
	}
	return o.Filters
}

// WithLimit set field Limit to given value
func (o *SearchOptions) WithLimit(value int) *SearchOptions {
	o.Limit = &value
	return o
}

// GetLimit returns value of field Limit
func (o *SearchOptions) GetLimit() int {
	if o.Limit == nil {
		var z int
		return z
	}
	return *o.Limit
}

// WithSkipTLSVerify set field SkipTLSVerify to given value
func (o *SearchOptions) WithSkipTLSVerify(value bool) *SearchOptions {
	o.SkipTLSVerify = &value
	return o
}

// GetSkipTLSVerify returns value of field SkipTLSVerify
func (o *SearchOptions) GetSkipTLSVerify() bool {
	if o.SkipTLSVerify == nil {
		var z bool
		return z
	}
	return *o.SkipTLSVerify
}

// WithListTags set field ListTags to given value
func (o *SearchOptions) WithListTags(value bool) *SearchOptions {
	o.ListTags = &value
	return o
}

// GetListTags returns value of field ListTags
func (o *SearchOptions) GetListTags() bool {
	if o.ListTags == nil {
		var z bool
		return z
	}
	return *o.ListTags
}
