package main

import (
	"errors"
	"fmt"
)

var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2-base
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2-base
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

type person struct {
	name string
	age  int
}

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
	var available bool // 一般声明
	valid := false     // 简短声明
	available = true   // 赋值操作
	fmt.Printf("Value is: %v, %v\n", valid, available)

	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v\n", c)

	m := `hello
	world`
	fmt.Printf("Value is: %v \n", m)

	err := errors.New("emit macho dwarf: elf header corrupted \n")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(slice)

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素

	//myFunc()
	//defterTest()
	var P person       // P现在就是person类型的变量了
	P.name = "Kaywall" // 赋值"Kaywall"给P的name属性.
	P.age = 25         // 赋值"25"给变量P的age属性
	fmt.Printf("The P's name is %s \n", P.name)

	P1 := person{"Tom1", 25}
	fmt.Printf("The P1's name is %s \n", P1.name)

	P2 := person{age: 24, name: "Tom2"}
	fmt.Printf("The P2's name is %s \n", P2.name)

	Point := new(person)
	fmt.Printf("The Point's name is %s \n", &Point.name)

	// 初始化学生Jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

func myFunc() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	if i > 10 && i > 0 {
		fmt.Println("end loop")
		println("test is ", i)
	} else {
		goto Here //跳转到Here去
	}
}

func defterTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
