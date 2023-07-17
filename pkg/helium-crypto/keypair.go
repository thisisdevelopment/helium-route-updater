package helium_crypto

import (
	"crypto/ed25519"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
)

/**
 * This is our golang implementation of https://github.com/helium/helium-crypto-rs/blob/main/src/keypair.rs
 *
 * For now we only support ED25519 keys
 *
 * The binary format for a keypair which contains both the public + private key is
 *   byte 0 => nettype | keytype
 *   byte 1-64 => private key (64 bytes, 512 bits)
 *   byte 65-97 => public key (32 bytes, 256 bits)
 *
 * Alternatively you can have a "keypair" which only has the public key; the binary format for that is:
 *   byte 0 => nettype | keytype
 *   byte 1-32 => public key (32 bytes, 256 bits)
 *
 * When the keys are displayed/used as string their keypairs are base58 encoded
 */

const (
	NETTYPE_MAIN byte = 0x00
	NETTYPE_TEST byte = 0x10

	KEYTYPE_SECP256K1   byte = 0x03
	KEYTYPE_MULTISIG    byte = 0x02
	KEYTYPE_ED25519     byte = 0x01
	KEYTYPE_ECC_COMPACT byte = 0x00
)

type KeyPair struct {
	NetType    byte
	KeyType    byte
	PublicKey  []byte
	PrivateKey []byte
}

func (k *KeyPair) Sign(message []byte) []byte {
	switch k.KeyType {
	case KEYTYPE_ED25519:
		return ed25519.Sign(k.PrivateKey, message)
	default:
		panic("not yet implemented")
	}
}

func (k *KeyPair) Verify(message []byte, signature []byte) bool {
	switch k.KeyType {
	case KEYTYPE_ED25519:
		return ed25519.Verify(k.PublicKey, message, signature)
	default:
		panic("not yet implemented")
	}
}

func (k *KeyPair) Public() *KeyPair {
	//TODO: cache this?
	return &KeyPair{
		NetType:    k.NetType,
		KeyType:    k.KeyType,
		PublicKey:  k.PublicKey,
		PrivateKey: nil,
	}
}

func (k *KeyPair) Bytes() []byte {
	res := []byte{k.NetType | k.KeyType}
	if k.PrivateKey == nil {
		res = append(res, k.PublicKey...)
	} else {
		// public key is contained in private key, so no need to output it separately here
		res = append(res, k.PrivateKey...)
	}
	return res
}

func (k *KeyPair) String() string {
	return base58.CheckEncode(k.Bytes(), 0)
}

func KeyPairFromBytes(input []byte) *KeyPair {
	var publicKey, privateKey []byte
	if len(input) == ed25519.PrivateKeySize+1 {
		privateKey = input[1:]
		publicKey = input[ed25519.PrivateKeySize-ed25519.PublicKeySize+1:]
	} else if len(input) == ed25519.PublicKeySize+1 {
		privateKey = nil
		publicKey = input[1:]
	} else {
		panic("invalid keypair length: " + strconv.FormatInt(int64(len(input)), 10))
	}

	return &KeyPair{
		NetType:    input[0] & 0xF0,
		KeyType:    input[0] & 0xF,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
}

func KeyPairFromString(input string) *KeyPair {
	bytes, _, _ := base58.CheckDecode(input)
	return KeyPairFromBytes(bytes)
}

func KeyPairGenerate(networkType byte, keyType byte) *KeyPair {
	switch keyType {
	case KEYTYPE_ED25519:
		publicKey, privateKey, _ := ed25519.GenerateKey(nil)
		return &KeyPair{
			NetType:    networkType,
			KeyType:    keyType,
			PublicKey:  publicKey,
			PrivateKey: privateKey,
		}
	default:
		panic("not yet implemented")
	}
}
