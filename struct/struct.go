package main

import (
	"fmt"
)

//Student 学生
type Student struct {
	name  string
	age   int
	score int
}

func main() {

	var jack Student
	jack.age = 18
	jack.name = "jack"
	jack.score = 92
	printStudent(jack)

	gjj := Student{name: "gjj", age: 23, score: 96}
	printStudent(gjj)

}

func printStudent(stu Student) {
	fmt.Println("==========", stu.name, "================")
	fmt.Printf("age\t%d\n", stu.age)
	fmt.Printf("score\t%d\n", stu.score)
	fmt.Println("--------------------------------")
}
