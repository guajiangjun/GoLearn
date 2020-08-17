package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//RSAGenKey ...
func RSAGenKey(lenBits int) error {

	//===============================私钥====================================
	//1.使用rsa的GenerateKey方法生成私钥，长度为lenBits
	privKey, err := rsa.GenerateKey(rand.Reader, lenBits)
	if err != nil {
		return err
	}
	//2.通过x509标准将得到的rsa私钥序列化为ASN.1的DER编码的字符串
	privStream := x509.MarshalPKCS1PrivateKey(privKey)
	//3.将私钥字符串设置到pem格式块中
	block := pem.Block{
		Type:  "RSA Private Key",
		Bytes: privStream,
	}
	//通过pem将设置好的数据进行编码，写入磁盘文件中
	privFile, err := os.Create("private.pem") //生成一个private.pem文件
	if err != nil {
		return err

	}
	defer privFile.Close()
	err = pem.Encode(privFile, &block) //将pem格式块写入private.pem文件中
	if err != nil {
		return err
	}
	//===============================公钥====================================
	pubKey := privKey.PublicKey
	//pubStream, err := x509.MarshalPKIXPublicKey(&pubKey)
	pubStream, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return nil
	}
	block = pem.Block{
		Type:  "RSA Public Key",
		Bytes: pubStream,
	}
	pubFile, err := os.Create("public.pem")
	if err != nil {
		return nil
	}
	err = pem.Encode(pubFile, &block)
	if err != nil {
		return nil
	}
	defer pubFile.Close()
	return nil
}
