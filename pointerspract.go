package main

import "fmt"

func main() {

	var x int = 5748
	var y int = 3456

	var p *int
	var q *int
	var s *int

	p = &x
	q = &y

	var sum1 = *p + *q
	//var sum2 = p + q
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in variable p = ", p)
	fmt.Println("Value stored in variable q = ", q)
	fmt.Println("Value stored in variable *q = ", *q)
	fmt.Println("s = ", s)
	fmt.Println(sum1)
	//fmt.Println(sum2)
}
