package validate

import (
	"github.com/hanks177/podman/v4/cmd/podman/registry"
	"github.com/spf13/cobra"
)

func AddLatestFlag(cmd *cobra.Command, b *bool) {
	// Initialization flag verification
	if !registry.IsRemote() {
		cmd.Flags().BoolVarP(b, "latest", "l", false,
			"Act on the latest container podman is aware of\nNot supported with the \"--remote\" flag")
	}
}
