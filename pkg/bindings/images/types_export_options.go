// Code generated by go generate; DO NOT EDIT.
package images

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *ExportOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *ExportOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithCompress set field Compress to given value
func (o *ExportOptions) WithCompress(value bool) *ExportOptions {
	o.Compress = &value
	return o
}

// GetCompress returns value of field Compress
func (o *ExportOptions) GetCompress() bool {
	if o.Compress == nil {
		var z bool
		return z
	}
	return *o.Compress
}

// WithFormat set field Format to given value
func (o *ExportOptions) WithFormat(value string) *ExportOptions {
	o.Format = &value
	return o
}

// GetFormat returns value of field Format
func (o *ExportOptions) GetFormat() string {
	if o.Format == nil {
		var z string
		return z
	}
	return *o.Format
}

// WithOciAcceptUncompressedLayers set field OciAcceptUncompressedLayers to given value
func (o *ExportOptions) WithOciAcceptUncompressedLayers(value bool) *ExportOptions {
	o.OciAcceptUncompressedLayers = &value
	return o
}

// GetOciAcceptUncompressedLayers returns value of field OciAcceptUncompressedLayers
func (o *ExportOptions) GetOciAcceptUncompressedLayers() bool {
	if o.OciAcceptUncompressedLayers == nil {
		var z bool
		return z
	}
	return *o.OciAcceptUncompressedLayers
}
