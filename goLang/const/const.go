package main

import "fmt"

const (
	BeiJing = iota * 10
	ShangHai
	ChongQing
)

func main() {
	const length int = 10
	fmt.Println(length)
	fmt.Println(BeiJing)
	fmt.Println(ShangHai)
	fmt.Println(ChongQing)
}
