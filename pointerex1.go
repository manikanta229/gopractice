package main

import "fmt"

func main() {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature) //it just store assigned value
	fmt.Println("pointer =", pointer)   //it stores the address of creature

	fmt.Println("*pointer =", *pointer) //it stores the value present in the address that is in pointer

	*pointer = "jellyfish" //we change the value so that is thje advantage of pointer without changing original variable
	fmt.Println("*pointer =", *pointer)
	fmt.Println("creature =", creature)

}
