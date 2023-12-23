package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
	age  int
}

func reflectNum(arg interface{}) {
	fmt.Println("Type: ", reflect.TypeOf(arg))
	fmt.Println("Value: ", reflect.ValueOf(arg))
}

func main18() {
	var num float64 = 3.1415926
	reflectNum(num)

	//var user User = User{1, "dbp", 19}
	interfaceMethod(num)
}

func interfaceMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Printf("InputType: %v, InputValue: %v\n", inputType, inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		t := inputType.Field(i)
		v := inputValue.Field(i)
		fmt.Println("Type: ", t.Name, t.Type, t.Anonymous, t.Index, t.Tag, t.Offset, t.PkgPath)
		fmt.Println("Value: ", v)
	}

}
