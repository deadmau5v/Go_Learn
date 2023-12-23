package main

//	多态
//
// Animal | interface 本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (_ *Cat) Sleep() {
	println("cat sleep...")
}

// GetColor 获取颜色
func (t *Cat) GetColor() string {
	return t.color
}

func (_ *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (_ *Dog) Sleep() {
	println("dog sleep...")
}

func (_ *Dog) GetColor() string {
	return "black"
}

func (_ *Dog) GetType() string {
	return "Dog"
}

func main() {
	// 多态
	var animal Animal
	animal = &Dog{}
	print(animal.GetType())
}
