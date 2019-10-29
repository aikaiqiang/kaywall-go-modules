package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	Name string
	Age  int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.Name + " - age: " + strconv.Itoa(p.Age) + " years)"
}

// 反射获取任意结构体类型和值（注意 大写字段表示公有）
func test(obj interface{}) {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i).Type
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		fmt.Printf("%d: %s %s = %v\n", i, fieldName, fieldType, fieldValue)
	}
}

func main() {
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 19}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("switch - list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("switch - list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("switch - list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("switch - list[%d] is of a different type", index)
		}
	}

	// 反射
	var x float32 = 3.4
	fmt.Println("x type:", reflect.TypeOf(x))
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)

	t := Person{"kaywall", 26}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	test(t)

}
