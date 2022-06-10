//go:build (amd64 && !windows) || (arm64 && !windows)
// +build amd64,!windows arm64,!windows

package machine

import (
	"github.com/hanks177/podman/v4/pkg/machine"
	"github.com/hanks177/podman/v4/pkg/machine/qemu"
)

func GetSystemDefaultProvider() machine.Provider {
	return qemu.GetQemuProvider()
}
