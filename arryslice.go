package main

import (
	"fmt"
)

func main() {
	var arry = [10]int{55, 99, 2, 11, 33, 77, 88, 2, 7, 1}

	for i := 0; i < len(arry); i++ {
		if arry[i]%2 == 0 {
			fmt.Println(arry[i])
		}

	}
	fmt.Println()
	a := arry[1:3]
	b := arry[:]
	fmt.Println(a)
	fmt.Println(b)
	a[1] = 3
	fmt.Println(a)

}
