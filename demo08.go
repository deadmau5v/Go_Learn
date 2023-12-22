package main

func main8() {
	// 字符串

	// 获取字符串长度
	var str1 string = "hello world"
	println(len(str1))

	// 截取
	println(str1[0:5])

	// 迭代 string
	for i := range str1[:5] {
		println(string(str1[i]))
	}

	// string 不能修改
	//str1[0] = 101 error
}
