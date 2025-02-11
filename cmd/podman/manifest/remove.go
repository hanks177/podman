package manifest

import (
	"fmt"

	"github.com/hanks177/podman/v4/cmd/podman/common"
	"github.com/hanks177/podman/v4/cmd/podman/registry"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:               "remove LIST IMAGE",
		Short:             "Remove an entry from a manifest list or image index",
		Long:              "Removes an image from a manifest list or image index.",
		RunE:              remove,
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: common.AutocompleteImages,
		Example:           `podman manifest remove mylist:v1.11 sha256:15352d97781ffdf357bf3459c037be3efac4133dc9070c2dce7eca7c05c3e736`,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: removeCmd,
		Parent:  manifestCmd,
	})
}

func remove(cmd *cobra.Command, args []string) error {
	updatedListID, err := registry.ImageEngine().ManifestRemoveDigest(registry.Context(), args[0], args[1])
	if err != nil {
		return errors.Wrapf(err, "error removing from manifest list %s", args[0])
	}
	fmt.Println(updatedListID)
	return nil
}
