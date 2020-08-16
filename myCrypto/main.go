package main

import (
	"fmt"
)

func main() {
	testDES()

}

func testDES() {

	fmt.Println("==========" + "DES 加解密" + "===========")
	str := `
作者：波罗学
链接：https://www.zhihu.com/question/327448992/answer/719337142
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。`
	src := []byte(str)
	fmt.Println("文本：\n", string(src))
	fmt.Println("文本编码成明文([]byte):\t", src, "\tlen:", len(src))
	key := []byte("abcdabcd")
	src2 := cryptDES(src, key)
	fmt.Println("加密后的密文([]byte)：\t", src2, "\tlen:", len(src2))
	src3 := decryptDES(src2, key)
	fmt.Println("密文解密后得到明文([]byte):\t", src3, "\tlen:", len(src3))
	fmt.Println("明文解码后得到文本：\n", string(src3))

}
