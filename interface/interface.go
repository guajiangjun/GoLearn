package main

import (
	"fmt"
)

//People 人 可以是男孩也可以是女孩
type People interface {
	like()
}

//Boy 男孩
type Boy struct {
	name string
	age  int
}

//Girl 女孩
type Girl struct {
}

func (boy Boy) like() {
	fmt.Println("我喜欢踢足球")
}

func (girl Girl) like() {
	fmt.Println("我喜欢画画")
}

func main() {

	var child People
	child = new(Boy)
	child.like()
	child = new(Girl)
	child.like()

}
