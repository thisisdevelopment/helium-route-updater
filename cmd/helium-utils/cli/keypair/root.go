package keypair

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
)

var RootCmd = &cobra.Command{
	Use:   "keypair",
	Short: "",
}

func init() {
	RootCmd.AddCommand(InfoCmd)
	RootCmd.AddCommand(GenerateCmd)
	RootCmd.AddCommand(PrivateCmd)
	RootCmd.AddCommand(PublicCmd)
	RootCmd.AddCommand(SignCmd)
	RootCmd.AddCommand(NewAllCmd(PublicCmd, PrivateCmd, InfoCmd))
}

func LoadKeyPair(input string) *helium_crypto.KeyPair {
	var keypair *helium_crypto.KeyPair
	if strings.HasPrefix(input, "/") || strings.HasPrefix(input, "./") {
		dat, err := os.ReadFile(input)
		if err != nil {
			log.Fatalf("Unable to read keypair from file: %s\n", input)
		}
		keypair = helium_crypto.KeyPairFromBytes(dat)
	} else {
		keypair = helium_crypto.KeyPairFromString(input)
	}
	return keypair
}
