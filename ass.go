package main

import (
	"ass1/pack"
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Println("enter value value of a,b and c: ")
	fmt.Scanln(&a, &b, &c)
	fmt.Println(pack.Add(a, b))
	fmt.Println(pack.Sun(b, c))
	fmt.Println(pack.Multi(a, c))
	fmt.Println(pack.Div(b, c))
	fmt.Println(pack.Squ(a))
	fmt.Println(pack.Pow(a))
	fmt.Println(pack.Asq(a, b))
	fmt.Println(pack.Nus(a, c))

}
