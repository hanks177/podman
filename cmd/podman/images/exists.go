package images

import (
	"github.com/hanks177/podman/v4/cmd/podman/common"
	"github.com/hanks177/podman/v4/cmd/podman/registry"
	"github.com/spf13/cobra"
)

var (
	existsCmd = &cobra.Command{
		Use:               "exists IMAGE",
		Short:             "Check if an image exists in local storage",
		Long:              `If the named image exists in local storage, podman image exists exits with 0, otherwise the exit code will be 1.`,
		Args:              cobra.ExactArgs(1),
		RunE:              exists,
		ValidArgsFunction: common.AutocompleteImages,
		Example: `podman image exists ID
  podman image exists IMAGE && podman pull IMAGE`,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: existsCmd,
		Parent:  imageCmd,
	})
}

func exists(cmd *cobra.Command, args []string) error {
	found, err := registry.ImageEngine().Exists(registry.GetContext(), args[0])
	if err != nil {
		return err
	}
	if !found.Value {
		registry.SetExitCode(1)
	}
	return nil
}
