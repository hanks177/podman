package entities

import (
	"net/url"

	"github.com/hanks177/podman/v4/libpod/define"
)

// VolumeCreateOptions provides details for creating volumes
// swagger:model
type VolumeCreateOptions struct {
	// New volume's name. Can be left blank
	Name string `schema:"name"`
	// Volume driver to use
	Driver string `schema:"driver"`
	// User-defined key/value metadata. Provided for compatibility
	Label map[string]string `schema:"label"`
	// User-defined key/value metadata. Preferred field, will override Label
	Labels map[string]string `schema:"labels"`
	// Mapping of driver options and values.
	Options map[string]string `schema:"opts"`
}

type VolumeConfigResponse struct {
	define.InspectVolumeData
}

type VolumeRmOptions struct {
	All     bool
	Force   bool
	Timeout *uint
}

type VolumeRmReport struct {
	Err error
	Id  string // nolint
}

type VolumeInspectReport struct {
	*VolumeConfigResponse
}

// VolumePruneOptions describes the options needed
// to prune a volume from the CLI
type VolumePruneOptions struct {
	Filters url.Values `json:"filters" schema:"filters"`
}

type VolumeListOptions struct {
	Filter map[string][]string
}

type VolumeListReport struct {
	VolumeConfigResponse
}

/*
 * Docker API compatibility types
 */

// VolumeMountReport describes the response from volume mount
type VolumeMountReport struct {
	Err  error
	Id   string // nolint
	Name string
	Path string
}

// VolumeUnmountReport describes the response from umounting a volume
type VolumeUnmountReport struct {
	Err error
	Id  string // nolint
}
