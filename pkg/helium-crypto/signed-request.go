package helium_crypto

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"reflect"
	"time"
)

type SignedRequest interface {
	proto.Message
	GetSignature() []byte
	GetSigner() []byte
	GetTimestamp() uint64
}

//TODO: as protoc-gen-go does not generate setters (https://github.com/golang/protobuf/issues/664) we revert to reflection here
//      the SignedRequest interface guarantees the fields are available

func SignRequest[T SignedRequest](req T, keypair *KeyPair) T {
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes([]byte{})
	reflect.ValueOf(req).Elem().FieldByName("Signer").SetBytes(keypair.Public().Bytes())
	//here to prevent replay attacks; although it does not seem to be checked
	reflect.ValueOf(req).Elem().FieldByName("Timestamp").SetUint(uint64(time.Now().UnixMilli()))
	marshalled, _ := proto.Marshal(req)
	signature := keypair.Sign(marshalled)
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes(signature)

	return req
}

func VerifyRequest[T SignedRequest](req T, keypair *KeyPair) bool {
	signature := req.GetSignature()
	signer := req.GetSigner()
	if bytes.Equal(keypair.Public().Bytes(), signer) {
		return false
	}
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes([]byte{})
	marshalled, _ := proto.Marshal(req)
	res := keypair.Verify(marshalled, signature)
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes(signature)

	return res
}
