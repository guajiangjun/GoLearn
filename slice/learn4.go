package main

import (
	"fmt"
)

func main() {

	s1 := []int{1, 2, 3, 4}

	fmt.Printf("type of s1 is:%T\n", s1)

	s2 := make([]int, 3, 8)
	s2 = []int{1, 2, 3}
	fmt.Println("s2:", s2)
	fmt.Printf("type of s2 is:%T，长度是:%d,容量是：%d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 2, 4)
	fmt.Println(s3)
	s3 = append(s3, s2...)
	fmt.Println(s3)

	fmt.Println("s2 is:", s2)
	fmt.Printf("s2.%%p-->%p\n", s2)
	fmt.Printf("&s2[0].%%p-->%p\n", &s2[0])
	fmt.Printf("&s2[1].%%p-->%p\n", &s2[1])
	fmt.Printf("&s2[2].%%p-->%p\n", &s2[2])

	ss := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss1 := ss[:]
	ss2 := ss[:2]
	ss3 := ss[3:9]
	fmt.Println("ss1 is :", ss1)
	fmt.Println("ss2 is :", ss2)
	fmt.Println("ss3 is :", ss3)

	fmt.Printf("address:ss-->%p\n", &ss)
	fmt.Printf("address:ss1-->%p\n", ss1)
	fmt.Printf("address:ss2-->%p\n", ss2)
	fmt.Printf("address:ss3-->%p\n", ss3)
	ss1[0] = 1
	fmt.Println(ss)

	sss := make([]int, 2, 10)
	fmt.Printf("cap is %d\nlen is %d\n", cap(sss), len(sss))

	sss = append(sss, 1, 2, 3, 4)
	fmt.Printf("cap is %d\nlen is %d\n", cap(sss), len(sss))

	sss = append(sss, 1, 2, 3, 4, 5, 6, 7, 7, 8, 8, 9, 3, 2, 3, 24)
	fmt.Printf("cap is %d\nlen is %d\n", cap(sss), len(sss))

	sss = append(sss, 1, 2, 3, 4, 5, 6, 7, 7, 8, 8, 9, 3, 2, 3, 24)
	fmt.Printf("cap is %d\nlen is %d\n", cap(sss), len(sss))
}
