package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
)

func main() {
	//input := "13WvV82S7QN3VMzMSieiGxvuaPKknMtf213E5JwPnboDkUfesKw"
	//input := "14MRZY2jc2ABDq1faCCMmXrkm2PXY9UBRTP1j9PWnFTKnCb7Hyn"
	//input := "66cDvko73yAf8LYvFMM3r8vF5vJtkk7JKMgEKwkmBC86oHdq41C7i1a2vS3zE1yCcdLLk6VUatUb32ZzVjSBXtRs"
	input := "F8UvVsKnzWyp2nF8aDcqvQ2GVcRpqT91WDsAtvBKCMt9"
	fmt.Printf("keypair: %+v\n", base58.Decode(input))
	b, _ := hex.DecodeString("018f23e96ab6bbff48c8923cac831dc97111bcf33dba9f5a8539c00f9d93551af1")
	keypair := helium_crypto.KeyPairFromBytes(b)

	signature, _ := hex.DecodeString("ef3e85dc7ea338c6b67399873131ea7b2265c516222e105fc39a59dda71f668a3b95fe27457d941a3cf5c422c9efbf0da112171d2997d74bc68f7b8118c6930e")
	message := "hello world"
	if keypair.Verify([]byte(message), signature) {
		fmt.Printf("This seems to work!\n")
	}
}
