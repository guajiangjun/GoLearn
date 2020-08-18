package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

//GetMd5Str1 ..
func GetMd5Str1(src []byte) {

	//md5.sum(): 源数据[]byte-->哈希值[16]byte
	res := md5.Sum(src)
	str := hex.EncodeToString(src)
	fmt.Println("len(src):", len(src))
	fmt.Println("len(str):", len(str))
	//[16]byte to string表示出来
	fmt.Println("res_string1:", hex.EncodeToString(res[:]))
}

//GetMd5Str2 ..
func GetMd5Str2(src []byte) {
	myHash := md5.New()
	//第一种添加数据的方式
	io.WriteString(myHash, "abcd")
	//第二种添加数据的方式
	myHash.Write([]byte("efg哈哈哈"))
	res := myHash.Sum(nil)

	fmt.Println("res_string2:", hex.EncodeToString(res))
}

// GetSHA256 ..
func GetSHA256(src []byte) {
	res := sha512.Sum512_256(src)
	fmt.Println("len(res):", len(res))
	str := hex.EncodeToString(src)
	fmt.Println("sha156 hash:", str)
	fmt.Println("len(str):", len(str))

}
