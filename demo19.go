package main

import "reflect"

type resume struct {
	Name   string `doc:"名字" info:"name"`
	Gender string `doc:"性别" info:"gender"`
}

func main19() {
	var res = resume{}
	findTag(res)
}

func findTag(arg interface{}) {
	inputType := reflect.TypeOf(arg)

	for i := 0; i < inputType.NumField(); i++ {
		t := inputType.Field(i)
		println(t.Tag)
		println(t.Tag.Get("info"))
	}
}
