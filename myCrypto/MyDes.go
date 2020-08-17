package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

//cryptDES 使用DES加密算法对明文进行加密-->{src:需要加密的明文,key:密钥,RETURN:加密后的密文}
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

//decryptDES 使用DES加密算法对密文进行解密-->{src:需要解密的明文,key:密钥,RETURN:解密后的密文}
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

func testDES() {

	str := `你好，我叫瓜江君，今年已经毕业了`
	fmt.Println("-----------文本：-----------------------------------\n", str)
	src := encode(str)
	fmt.Println("\n======================" + "DES 加解密" + "=======================================================================")
	fmt.Println("[明文]：\n", src, "\nlen:", len(src))
	key := []byte("abcdabcd")
	src2 := cryptDES(src, key)
	fmt.Println("[密文]：\n", src2, "\nlen:", len(src2))
	src3 := decryptDES(src2, key)
	fmt.Println("[明文]\n", src3, "\nlen:", len(src3))
	fmt.Println("======================" + "DES 加解密 END" + "=======================================================================")
	fmt.Println("\n-----------明文解码后得到文本：--------------------------\n", decode(src3))

}
