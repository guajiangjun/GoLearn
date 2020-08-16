package main

import "bytes"

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
