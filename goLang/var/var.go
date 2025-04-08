package main

import "fmt"

// 声明变量，不初始化默认为0
func main() {
	var a int = 0
	fmt.Printf("a=%d\n", a)

	var c = 100
	d := 200
	fmt.Printf("c=%d\n", c)
	fmt.Printf("a=%T\nc=%T\n", a, c)

	fmt.Printf("d=%d\n", d)
	var bb = "golang"
	fmt.Printf("bb=%s\n", bb)
}
