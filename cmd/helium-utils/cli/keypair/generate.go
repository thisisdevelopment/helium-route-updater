package keypair

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
)

type generateConfig struct {
	netType string
	keyType string
}

var GenerateCmd *cobra.Command
var config = generateConfig{
	keyType: "ed25519",
}

var rootCmd = &cobra.Command{
	Use: "generate OUTPUT_PATH",
	Short: "Generate new KeyPair",
	Long: "Generates a new KeyPair file with embedded Net and Key Types.",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		netType, err := parseArgNetType(config.netType)
		if err != nil {
			return err
		}
		
		keyType, err := parseArgKeyType(config.keyType)
		if err != nil {
			return err
		}
		
		kp := helium_crypto.KeyPairGenerate(netType, keyType)

		err = os.WriteFile(args[0], kp.Bytes(), 0644)
		if err != nil {
			return err
		}

		fmt.Println("keypair written to " + args[0])
		return nil
	},
}

func init() {
	GenerateCmd = rootCmd
	
	rootCmd.Flags().StringVarP(
		&config.netType,
		"net_type",
		"n",
		"",
		"net type, valid options: (required)\n - main\n - test")
	rootCmd.MarkFlagRequired("net_type")

	rootCmd.Flags().StringVarP(
		&config.keyType,
		"key_type",
		"k",
		"ed25519",
		"key type, valid options:\n - 2d25519")
}

func parseArgNetType(arg string) (byte, error) {
	switch strings.ToLower(arg) {
	case "main":
		return helium_crypto.NETTYPE_MAIN, nil
	case "test":
		return helium_crypto.NETTYPE_TEST, nil
	default:
		return 0, errors.New("invalid net type. " + arg + " given, must be one of main,test.")
	}
}
func parseArgKeyType(arg string) (byte, error) {
	switch strings.ToLower(arg) {
	case "ed25519":
		return helium_crypto.KEYTYPE_ED25519, nil
	default:
		return 0, errors.New("invalid key type. " + arg + " given, must be ed25519")
	}
}
