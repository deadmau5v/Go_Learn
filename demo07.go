package main

func main7() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			print(i, " X ", j, " = ", i*j, "|")
		}
		println()
	}
}
