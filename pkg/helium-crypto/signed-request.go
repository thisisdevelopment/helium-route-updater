package helium_crypto

import (
	"google.golang.org/protobuf/proto"
	"reflect"
)

type SignedRequest interface {
	proto.Message
	GetSignature() []byte
}

//TODO: as protoc-gen-go does not generate setters (https://github.com/golang/protobuf/issues/664) we revert to reflection here

func SignRequest[T SignedRequest](req T, keypair *KeyPair) T {
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes([]byte{})
	marshalled, _ := proto.Marshal(req)
	signature := keypair.Sign(marshalled)
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes(signature)

	return req
}

func VerifyRequest[T SignedRequest](req T, keypair *KeyPair) bool {
	signature := req.GetSignature()
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes([]byte{})
	marshalled, _ := proto.Marshal(req)
	res := keypair.Verify(marshalled, signature)
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes(signature)

	return res
}
