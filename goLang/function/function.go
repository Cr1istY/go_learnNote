package main

import "fmt"

func FunOne(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c := 10
	return c
}

func main() {
	var c = 0
	c = FunOne("hello world", 10)
	fmt.Println(c)
}
