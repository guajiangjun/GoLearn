package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func cryptAES(src, key []byte) []byte {

	//生成加密函数
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCEncrypter(block, key)

	//将明文进行padding
	dst := padding(src, block.BlockSize())

	//将明文加密成密文
	blockMode.CryptBlocks(dst, dst)
	return dst
}

func decryptAES(src, key []byte) []byte {

	//生成解密函数
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key)

	//将密文加密得到明文
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)

	//将明文解pdding
	dst = unPadding(dst)
	return dst

}

func testAES(str string) {

	key := []byte("abcdabcd12345678")

	fmt.Println("|===============================================|")
	fmt.Println("|============= myAES 加解密 测试 ===============|")
	fmt.Println("|===============================================|")

	fmt.Println("文本：", str)

	src := encode(str)
	fmt.Printf("明文：（len: %d）:\n%v\n", len(src), src)

	src = cryptAES(src, key)
	fmt.Printf("密文：（len: %d）:\n%v\n", len(src), src)

	src = decryptAES(src, key)
	fmt.Printf("明文：（len: %d）:\n%v\n", len(src), src)
	fmt.Println("文本：", str)

}
