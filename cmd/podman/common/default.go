package common

import (
	"github.com/hanks177/podman/v4/cmd/podman/registry"
)

var (
	// Pull in configured json library
	json = registry.JSONLibrary()
)
