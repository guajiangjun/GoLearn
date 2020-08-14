package main

import (
	"fmt"
)

func main() {
	sum := getsum(100)
	defer fmt.Println("the sum is:", sum)

	defer fmt.Println("...main...")
	defer sayhello()

}

func sayhello() {
	fmt.Println("sayhello..function...")
}

//getsum 计算1+2+3+...+n之和
func getsum(n int) int {
	if n == 1 {
		return 1
	}
	return getsum(n-1) + n

}
