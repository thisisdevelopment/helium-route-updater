package check

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := types.ConfigFromEnv()

		if config.Helium.Server == "" {
			return errors.New("Helium.Server not set")
		}

		if config.Helium.RouteId == "" {
			return errors.New("Helium.RouteId not set")
		}

		if config.Helium.KeyPair == "" {
			return errors.New("Helium.KeyPair not set")
		}

		if config.Lns.Type == "" {
			return errors.New("Lns.Type not set")
		}

		if config.Lns.Listen == "" {
			return errors.New("Lns.Listen not set")
		}

		if config.Lns.ApiAuth == "" {
			return errors.New("Lns.ApiAuth not set")
		}

		if config.Lns.ApiEndpoint == "" {
			return errors.New("Lns.ApiEndpoint not set")
		}

		// everything OK, don't return error
		return nil
	},
}
