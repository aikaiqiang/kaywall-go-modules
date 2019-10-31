package main

import (
	"fmt"
	"strconv"
)

type Human2 struct {
	name  string
	age   int
	phone string
}

type Human struct {
	name   string
	age    int
	weight int
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human2) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

// Human 实现了 error 接口的对象
func (h Human) Error() string {
	return "" + h.name + ""
}

func main() {
	Bob := Human2{"Bob", 39, "000-7777-XXX"}
	Jack := Human{"Jack", 26, 120}
	fmt.Println("This Human is : ", Bob)
	fmt.Println("This Human is : ", Jack)
}
