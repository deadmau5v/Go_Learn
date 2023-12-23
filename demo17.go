package main

import (
	"fmt"
	"reflect"
)

type master interface {
	Say()
}

type slaver interface {
	Do()
}

type Student1 struct {
	name string
	age  int
}

func (s *Student1) Say() {
	println(s.name, ": Just Do It !")
}

func (s *Student1) Do() {
	println(s.name, ": Do it...")

}

func main17() {
	student := Student1{"dbp", 18}
	student2 := Student1{"lz", 18}

	var m master
	m = &student
	m.Say()

	var s slaver
	s = &student2
	s.Do()

	valueAndType(student)
	valueAndType(&student2)
}

func valueAndType(any interface{}) {
	fmt.Printf(
		"<%v, %v>\n",
		reflect.TypeOf(any),
		reflect.ValueOf(any),
	)
}
