package cli

import (
	"github.com/spf13/cobra"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/check"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/helium"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/keypair"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/lns"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "helium-utils",
	Short: "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		c := types.ConfigFromEnv()
		cmd.SetContext(config.WithContext(nil, c))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(check.RootCmd)
	RootCmd.AddCommand(lns.RootCmd)
	RootCmd.AddCommand(helium.RootCmd)
	RootCmd.AddCommand(keypair.RootCmd)
	RootCmd.AddCommand(syncCmd)
	RootCmd.AddCommand(infoCmd)
}
