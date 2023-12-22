package main

import (
	"fmt"
)

func main5() {
	println("启动")
	var source int
	fmt.Scanln(&source)

	if source >= 95 {
		println("A+")
	} else if 95 > source && source >= 90 {
		println("A")
	} else if 90 > source && source >= 80 {
		println("B")
	} else if 80 > source && source >= 70 {
		println("C+")
	} else if 70 > source && source >= 60 {
		println("C")
	} else {
		println("D")
	}

	// switch 语句
	switch source {
	case 100, 99:
		println("优秀!")
	case 60, 61, 62:
		println("擦边？")
	case 57, 58, 59:
		println("笑死我了")
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		println("纯纯傻逼!")
	default:
		println("一般吧")
	}

	switch source {
	case 100:
		println("牛逼")
	default:
		println("这位更是重量级")
	}

	var sum int = 0

	for i := 1; i <= 1000; i++ {
		println(i)
		sum += i
	}

	println(sum)

	var i int = 0
	for ; ; println(1) {
		i++
		println(0)
		if i == 1000 {
			println("循环了1k次")
			break
		}
	}

	for {
		println("无限循环")
	}

}
