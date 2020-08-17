package main

import "fmt"

func main() {
	RSATest()
}

//RSATest ...
func RSATest() {
	err := RSAGenKey(1024)
	if err != nil {
		fmt.Println("错误信息：", err)
		return
	}
	src := []byte("所以你叫什么名字，但是你到底是想干什么？")
	fmt.Println("明文：" + string(src))
	msg, err := RSAPublicEncrypt(src, "public.pem")
	if err != nil {
		fmt.Println("公钥加密失败：", err)
		return
	}
	fmt.Println("len of src is :", len(src))
	fmt.Println("len of msg is :", len(msg))
	src, err = RSAPrivateDecrypt(msg, "private.pem")

	if err != nil {
		fmt.Println("私钥解密失败：", err)
		return
	}
	fmt.Println("len of src is :", len(src))
	fmt.Println("解密得到：" + string(src))
	return

}
