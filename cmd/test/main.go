package main

import (
	"encoding/hex"
	"fmt"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
)

func main() {
	input := "13WvV82S7QN3VMzMSieiGxvuaPKknMtf213E5JwPnboDkUfesKw"
	keypair := helium_crypto.KeyPairFromString(input)

	signature, _ := hex.DecodeString("ef3e85dc7ea338c6b67399873131ea7b2265c516222e105fc39a59dda71f668a3b95fe27457d941a3cf5c422c9efbf0da112171d2997d74bc68f7b8118c6930e")
	message := "hello world"
	if keypair.Verify([]byte(message), signature) {
		fmt.Printf("This seems to work!\n")
	}

	kp := helium_crypto.KeyPairGenerate(helium_crypto.NETTYPE_MAIN, helium_crypto.KEYTYPE_ED25519)
	fmt.Printf("Public key: %s\n", kp.Public().String())
	fmt.Printf("Private key: %s\n", hex.EncodeToString(kp.Bytes()))
}
