package connection

import (
	"fmt"

	"github.com/containers/common/pkg/config"
	"github.com/hanks177/podman/v4/cmd/podman/common"
	"github.com/hanks177/podman/v4/cmd/podman/registry"
	"github.com/hanks177/podman/v4/cmd/podman/system"
	"github.com/spf13/cobra"
)

var (
	// Skip creating engines since this command will obtain connection information to said engines.
	dfltCmd = &cobra.Command{
		Use:               "default NAME",
		Args:              cobra.ExactArgs(1),
		Short:             "Set named destination as default",
		Long:              `Set named destination as default for the Podman service`,
		ValidArgsFunction: common.AutocompleteSystemConnections,
		RunE:              defaultRunE,
		Example:           `podman system connection default testing`,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: dfltCmd,
		Parent:  system.ConnectionCmd,
	})
}

func defaultRunE(cmd *cobra.Command, args []string) error {
	cfg, err := config.ReadCustomConfig()
	if err != nil {
		return err
	}

	if _, found := cfg.Engine.ServiceDestinations[args[0]]; !found {
		return fmt.Errorf("%q destination is not defined. See \"podman system connection add ...\" to create a connection", args[0])
	}

	cfg.Engine.ActiveService = args[0]
	return cfg.Write()
}
