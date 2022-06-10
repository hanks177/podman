package qemu

import (
	"github.com/hanks177/podman/v4/pkg/rootless"
	"github.com/hanks177/podman/v4/pkg/util"
)

func getRuntimeDir() (string, error) {
	if !rootless.IsRootless() {
		return "/run", nil
	}
	return util.GetRuntimeDir()
}
