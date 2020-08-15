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
	changeAge(jack, 11)
	fmt.Println("+++++++++after changename+++++++++++++")
	printStudent(jack)
	changeAgeByPointer(&jack, 11)
	fmt.Println("+++++++++after changenamebypointer+++++++++++++")
	printStudent(jack)

}

func printStudent(stu Student) {
	fmt.Println("==========", stu.name, "================")
	fmt.Printf("age\t%d\n", stu.age)
	fmt.Printf("score\t%d\n", stu.score)
	fmt.Println("--------------------------------")
}

func printStudentByPointer(stu *Student) {
	fmt.Println("==========", stu.name, "================")
	fmt.Printf("age\t%d\n", stu.age)
	fmt.Printf("score\t%d\n", stu.score)
	fmt.Println("--------------------------------")
}

func changeAge(stu Student, age int) {
	stu.age = age
}

func changeAgeByPointer(stu *Student, age int) {
	stu.age = age
}
