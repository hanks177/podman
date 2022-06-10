// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/hanks177/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *RestoreOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *RestoreOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithIgnoreRootfs set field IgnoreRootfs to given value
func (o *RestoreOptions) WithIgnoreRootfs(value bool) *RestoreOptions {
	o.IgnoreRootfs = &value
	return o
}

// GetIgnoreRootfs returns value of field IgnoreRootfs
func (o *RestoreOptions) GetIgnoreRootfs() bool {
	if o.IgnoreRootfs == nil {
		var z bool
		return z
	}
	return *o.IgnoreRootfs
}

// WithIgnoreVolumes set field IgnoreVolumes to given value
func (o *RestoreOptions) WithIgnoreVolumes(value bool) *RestoreOptions {
	o.IgnoreVolumes = &value
	return o
}

// GetIgnoreVolumes returns value of field IgnoreVolumes
func (o *RestoreOptions) GetIgnoreVolumes() bool {
	if o.IgnoreVolumes == nil {
		var z bool
		return z
	}
	return *o.IgnoreVolumes
}

// WithIgnoreStaticIP set field IgnoreStaticIP to given value
func (o *RestoreOptions) WithIgnoreStaticIP(value bool) *RestoreOptions {
	o.IgnoreStaticIP = &value
	return o
}

// GetIgnoreStaticIP returns value of field IgnoreStaticIP
func (o *RestoreOptions) GetIgnoreStaticIP() bool {
	if o.IgnoreStaticIP == nil {
		var z bool
		return z
	}
	return *o.IgnoreStaticIP
}

// WithIgnoreStaticMAC set field IgnoreStaticMAC to given value
func (o *RestoreOptions) WithIgnoreStaticMAC(value bool) *RestoreOptions {
	o.IgnoreStaticMAC = &value
	return o
}

// GetIgnoreStaticMAC returns value of field IgnoreStaticMAC
func (o *RestoreOptions) GetIgnoreStaticMAC() bool {
	if o.IgnoreStaticMAC == nil {
		var z bool
		return z
	}
	return *o.IgnoreStaticMAC
}

// WithImportAchive set field ImportAchive to given value
func (o *RestoreOptions) WithImportAchive(value string) *RestoreOptions {
	o.ImportAchive = &value
	return o
}

// GetImportAchive returns value of field ImportAchive
func (o *RestoreOptions) GetImportAchive() string {
	if o.ImportAchive == nil {
		var z string
		return z
	}
	return *o.ImportAchive
}

// WithImportArchive set field ImportArchive to given value
func (o *RestoreOptions) WithImportArchive(value string) *RestoreOptions {
	o.ImportArchive = &value
	return o
}

// GetImportArchive returns value of field ImportArchive
func (o *RestoreOptions) GetImportArchive() string {
	if o.ImportArchive == nil {
		var z string
		return z
	}
	return *o.ImportArchive
}

// WithKeep set field Keep to given value
func (o *RestoreOptions) WithKeep(value bool) *RestoreOptions {
	o.Keep = &value
	return o
}

// GetKeep returns value of field Keep
func (o *RestoreOptions) GetKeep() bool {
	if o.Keep == nil {
		var z bool
		return z
	}
	return *o.Keep
}

// WithName set field Name to given value
func (o *RestoreOptions) WithName(value string) *RestoreOptions {
	o.Name = &value
	return o
}

// GetName returns value of field Name
func (o *RestoreOptions) GetName() string {
	if o.Name == nil {
		var z string
		return z
	}
	return *o.Name
}

// WithTCPEstablished set field TCPEstablished to given value
func (o *RestoreOptions) WithTCPEstablished(value bool) *RestoreOptions {
	o.TCPEstablished = &value
	return o
}

// GetTCPEstablished returns value of field TCPEstablished
func (o *RestoreOptions) GetTCPEstablished() bool {
	if o.TCPEstablished == nil {
		var z bool
		return z
	}
	return *o.TCPEstablished
}

// WithPod set field Pod to given value
func (o *RestoreOptions) WithPod(value string) *RestoreOptions {
	o.Pod = &value
	return o
}

// GetPod returns value of field Pod
func (o *RestoreOptions) GetPod() string {
	if o.Pod == nil {
		var z string
		return z
	}
	return *o.Pod
}

// WithPrintStats set field PrintStats to given value
func (o *RestoreOptions) WithPrintStats(value bool) *RestoreOptions {
	o.PrintStats = &value
	return o
}

// GetPrintStats returns value of field PrintStats
func (o *RestoreOptions) GetPrintStats() bool {
	if o.PrintStats == nil {
		var z bool
		return z
	}
	return *o.PrintStats
}

// WithPublishPorts set field PublishPorts to given value
func (o *RestoreOptions) WithPublishPorts(value []string) *RestoreOptions {
	o.PublishPorts = value
	return o
}

// GetPublishPorts returns value of field PublishPorts
func (o *RestoreOptions) GetPublishPorts() []string {
	if o.PublishPorts == nil {
		var z []string
		return z
	}
	return o.PublishPorts
}

// WithFileLocks set field FileLocks to given value
func (o *RestoreOptions) WithFileLocks(value bool) *RestoreOptions {
	o.FileLocks = &value
	return o
}

// GetFileLocks returns value of field FileLocks
func (o *RestoreOptions) GetFileLocks() bool {
	if o.FileLocks == nil {
		var z bool
		return z
	}
	return *o.FileLocks
}
