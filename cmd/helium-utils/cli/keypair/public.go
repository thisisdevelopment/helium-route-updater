package keypair

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

var PublicCmd = &cobra.Command{
	Use:   "public",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := types.ConfigFromEnv()

		keypair := LoadKeyPair(config.Helium.KeyPair)

		if keypair.PublicKey == nil {
			return errors.New("failed to load public key from " + config.Helium.KeyPair)
		}

		fmt.Println("Public Keypair OK")
		return nil
	},
}
