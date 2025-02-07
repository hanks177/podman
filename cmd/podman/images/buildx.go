package images

import (
	"github.com/hanks177/podman/v4/cmd/podman/registry"
	"github.com/hanks177/podman/v4/cmd/podman/validate"
	"github.com/spf13/cobra"
)

var (
	// Command: podman _buildx_
	// This is a hidden command, which was added to make converting
	// from Docker to Podman easier.
	// For now podman buildx build just calls into podman build
	// If we are adding new buildx features, we will add them by default
	// to podman build.
	buildxCmd = &cobra.Command{
		Use:     "buildx",
		Aliases: []string{"builder"},
		Short:   "Build images",
		Long:    "Build images",
		RunE:    validate.SubCommandExists,
		Hidden:  true,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: buildxCmd,
	})
}
