package main

import (
	"crypto/cipher"
	"crypto/des"
)

//crypt3DES 3DES加密函数-->{src:明文，key:密钥，RETURN:密文}
func crypt3DES(src, key []byte) []byte {

	//生成加密函数
	block, err := des.NewTripleDESCipher(key) //密钥
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCEncrypter(block, key[0:block.BlockSize()]) //初始向量

	//对明文进行填充
	dst := padding(src, block.BlockSize())

	//对填充后的明文进行加密
	blockMode.CryptBlocks(dst, dst)
	return dst

}

//decrypt3DES 3DES解密函数-->{src:密文，key:密钥，RETURN:明文}
func decrypt3DES(src, key []byte) []byte {

	//生成解密函数
	block, err := des.NewTripleDESCipher(key) //密钥
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()]) //初始向量

	//将密文解密成明文
	blockMode.CryptBlocks(src, src)

	//将明文解padding
	src = unPadding(src)
	return src

}
