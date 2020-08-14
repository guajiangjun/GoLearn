package main

import "fmt"

func main() {
	{ //变量定义的三种方法
		//第一种
		var num1 int
		num1 = 30
		fmt.Println("=================第1种定义======================")
		fmt.Printf("num1=%d\n", num1)
		var num2 int = 56
		fmt.Printf("num2=%d\n", num2)
		fmt.Println()
		//第二种
		fmt.Println("=================第2种定义======================")
		var num3 = "3333"
		fmt.Printf("num3的类型是:%T,值是:%s\n", num3, num3)
		fmt.Println()
		//第三种
		fmt.Println("=================第3种定义======================")
		num4 := 3.14156
		fmt.Printf("num4的类型是:%T,值是:%f\n", num4, num4)
		fmt.Println()
		//多个变量同时定义
		fmt.Println("=================多个变量同时定义================")
		var a, b, c int
		a = 1
		b = 2
		c = 3
		fmt.Println(a, b, c)

		var n1, f1, s1 = 100, 2.71676, "GJJ"
		fmt.Println(n1, f1, s1)

		var (
			name = "GJJ"
			age  = 28
			high = 178.7676
			sex  = "男"
		)
		fmt.Printf("我的姓名是:%s,我的年龄是:%d,我的身高是:%fcm,我的性别是:%s\n", name, age, high, sex)

		myname, myage := "比克", 26
		fmt.Printf("我的姓名是:%s,我的年龄是:%d\n", myname, myage)

	}
	var num int = 3
	fmt.Println("你好，我今年", num, "岁了！")
	for i := 1; i <= 5; i++ {
		fmt.Println("i=", i)
	}

	fmt.Println("====================================")
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, "\t")
			count++
			if count%5 == 0 {
				fmt.Print("\n")
			}

		}
	}
	fmt.Print("\n")
	fmt.Println("count-->", count)
	fmt.Println("====================================")
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d X %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	var arr1 = [4]int{11, 13, 17, 88}
	for index, value := range arr1 {
		fmt.Println("下标是:", index, ",值是:", value)
	}
	for _, value := range arr1 {
		fmt.Println("value:", value)
	}
	fmt.Println("====================================")
	add(2, 3, 3)
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	add(s1...)
	fmt.Println("====================================")
	w1 := [4]int{1, 2, 3, 4}
	w2 := []int{2, 6, 8, 9}

	fmt.Printf("type of w1 is %T\n", w1)
	fmt.Printf("type of w2 is %T\n", w2)

}

func add(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Printf("the result is:%d\n", sum)
}
