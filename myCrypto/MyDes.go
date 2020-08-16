package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

//PADDING 密钥填充函数
/*
[srctext:需要解填充原始文本,
blocksize:密钥分组长度,
RETURN:填充后的密钥]
*/
func padding(srctext []byte, blocksize int) []byte {

	paddingSize := blocksize - len(srctext)%blocksize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	paddedText := append(srctext, paddingText...)
	return paddedText
}

//UNPADDING 密钥解填充函数
/*
[srcText:需要解填充原始文本,
RETURN:解填充后的文本]
*/
func unPadding(srcText []byte) []byte {

	len := len(srcText)
	paddingSize := int(srcText[len-1])

	unPaddedText := srcText[:len-paddingSize]
	return unPaddedText

}

//cryptDES 使用DES加密算法对明文进行加密
/*
[src:需要加密的明文
key:密钥
RETURN:加密后的密文]
*/
func cryptDES(src, key []byte) []byte {

	//(声明密钥)创建并返回一个，密钥为key的，使用DES加密的，cipher.block接口
	block, _ := des.NewCipher(key)
	//（声明初始向量）生成一个和密钥等长的初始“向量”
	iv := []byte("abcdabcd")
	//（创将加密函数）创建一个CBC模式的，底层为DES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, iv)

	src = padding(src, block.BlockSize()) //对明文进行padding
	dst := make([]byte, len(src))

	//加密连续的数据块
	blockMode.CryptBlocks(dst, src)
	return dst

}

//decryptDES 使用DES加密算法对密文进行解密
/*
[src:需要解密的明文
key:密钥
RETURN:解密后的密文]
*/
func decryptDES(src, key []byte) []byte {

	//(声明密钥)创建并返回一个，密钥为key的，使用DES加密的，cipher.block接口
	block, _ := des.NewCipher(key)
	//（声明初始向量）生成一个和密钥等长的初始“向量”
	iv := []byte("abcdabcd")
	//（创将解密函数）创建一个CBC模式的，底层为DES加密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, iv)

	dst := make([]byte, len(src))

	//解密密文
	blockMode.CryptBlocks(dst, src)

	dst = unPadding(dst)
	return dst

}
