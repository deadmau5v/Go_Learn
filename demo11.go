package main

func main11() {
	test1()

	var a int = 1
	var b *int = &a
	println("b 的内存地址", b)
	var c **int = &b
	println("c 的内存地址", c)
	var cc *int = *c
	println("**c 的值", *cc) // 二级指针

}

func test1() {
	defer println("...")
	defer println("作用域 test")
	defer println("正在推出 test1")
	println("执行test1")
}
