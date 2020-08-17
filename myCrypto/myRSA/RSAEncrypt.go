package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//RSAPublicEncrypt ...
func RSAPublicEncrypt(src []byte, pathname string) ([]byte, error) {
	msg := []byte("")
	file, err := os.Open(pathname)
	if err != nil {
		return msg, err
	}
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}
	recvBuff := make([]byte, info.Size())
	file.Read(recvBuff)
	block, _ := pem.Decode(recvBuff)
	pubInter, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	pubKey := pubInter.(*rsa.PublicKey)

	msg, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

//RSAPrivateDecrypt ...
func RSAPrivateDecrypt(src []byte, filename string) ([]byte, error) {

	msg := []byte("")
	file, err := os.Open(filename)
	if err != nil {
		return msg, err
	}
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}
	recvBuff := make([]byte, info.Size())
	file.Read(recvBuff)
	block, _ := pem.Decode(recvBuff)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	msg, err = rsa.DecryptPKCS1v15(rand.Reader, privKey, src)
	if err != nil {
		return msg, nil
	}
	return msg, nil
}
