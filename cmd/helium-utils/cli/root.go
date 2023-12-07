package cli

import (
	"github.com/spf13/cobra"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/check"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/helium"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/keypair"
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli/lns"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "helium-utils",
	Short: "",
}

func init() {
	RootCmd.AddCommand(check.RootCmd)
	RootCmd.AddCommand(lns.RootCmd)
	RootCmd.AddCommand(helium.RootCmd)
	RootCmd.AddCommand(keypair.RootCmd)
	//RootCmd.AddCommand(syncCmd)
	//RootCmd.AddCommand(infoCmd)
}
