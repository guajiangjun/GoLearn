package main

import (
	"fmt"
)

func main() {
	test3DES()

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

func test3DES() {

	str := `今天是一个晴朗的天气`
	fmt.Println("-----------文本：-----------------------------------\n", str)
	src := encode(str)
	fmt.Println("\n======================" + "DES 加解密" + "=======================================================================")
	fmt.Println("[明文]：\n", src, "\n明文长度:", len(src))
	key := []byte("abcdefghabcdefghabcdefgh")
	src2 := crypt3DES(src, key)
	fmt.Println("[密文]：\n", src2, "\n密文长度:", len(src2))
	src3 := decrypt3DES(src2, key)
	fmt.Println("[明文]\n", src3, "\n明文长度:", len(src3))
	fmt.Println("======================" + "DES 加解密 END" + "=======================================================================")
	fmt.Println("\n-----------明文解码后得到文本：--------------------------\n", decode(src3))

}
