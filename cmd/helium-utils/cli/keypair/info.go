package keypair

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		config := types.ConfigFromEnv()

		keypair := LoadKeyPair(config.Helium.KeyPair)


		if nil == keypair {
			log.Fatal(errors.New("failed to load keypair"))
		}

		var t string; switch keypair.NetType {
		case helium_crypto.NETTYPE_MAIN:
			t = "MAIN"
			break
		case helium_crypto.NETTYPE_TEST:
			t = "TEST"
			break
		default:
			t = "UNKNOWN"
		}

		var c string; switch keypair.KeyType {
		case helium_crypto.KEYTYPE_ED25519:
			c = "ed25519"
			break
		case helium_crypto.KEYTYPE_ECC_COMPACT:
			c = "ecc-compact"
			break
		case helium_crypto.KEYTYPE_SECP256K1:
			c = "secp256k1"
			break
		case helium_crypto.KEYTYPE_MULTISIG:
			c = "multisig"
			break
		default:
			c = "unknown"
		}

		fmt.Println("Net Type: " + t)
		fmt.Println("Key Type: " + c)

		fmt.Println(fmt.Sprintf("Private Key (%d): %X", len(keypair.PrivateKey), keypair.PrivateKey))
		fmt.Println(fmt.Sprintf("Public key (%d): %X", len(keypair.PublicKey), keypair.PublicKey))
	},
}

