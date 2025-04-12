package main

import "fmt"

func main() {
	var arr [10]int
	myarr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}

	for i := 0; i < len(myarr); i++ {
		fmt.Println(myarr[i])
	}

	for index, value := range arr {
		fmt.Println("index = ", index, "value = ", value)
	}
}
