package main

import "fmt"

func main() {
	s1 := [2]int{10, 10}
	s := []int{1, 2, 3, 4}
	fmt.Println("the sum is:", add(1, 2, s1, s...))

	fmt.Println("=============================================")
	L, S := rectangle(2, 3)
	fmt.Println("面积:", S, "周长:", L)
	fmt.Println("=============================================")
	L1, _ := rectangle(1, 1)
	fmt.Println("周长:", L1)

}

//列表，切片作为参数
func add(a, b int, c [2]int, nums ...int) (sum int) {
	sum += a + b
	for i := 0; i < len(c); i++ {
		sum += c[i]
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return
}

//返回多个值
func rectangle(length, width float64) (L, S float64) {
	L = 2 * (length + width)
	S = length * width
	return
}
