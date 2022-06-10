package machine

import (
	"github.com/hanks177/podman/v4/pkg/machine"
	"github.com/hanks177/podman/v4/pkg/machine/wsl"
)

func GetSystemDefaultProvider() machine.Provider {
	return wsl.GetWSLProvider()
}
