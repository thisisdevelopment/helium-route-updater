package keypair

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

var PrivateCmd = &cobra.Command{
	Use:   "private",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := types.ConfigFromEnv()

		keypair := LoadKeyPair(config.Helium.KeyPair)

		if keypair.PrivateKey == nil {
			return errors.New("failed to load private key from " + config.Helium.KeyPair)
		}

		fmt.Println("Private Keypair OK")
		return nil
	},
}
