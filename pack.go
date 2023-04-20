package pack

import "math"

func Add(a, b int) int {
	return a + b
}
func Sun(a, b int) int {
	return a - b
}
func Multi(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	return a //b
}
func Squ(a int) int {
	return a * a
}
func Pow(a int) int {
	return int(math.Pow10(a))

}
func Asq(a, b int) int {
	return ((a * a) + (b * b) + 2*(a*b))
}
func Nus(a, b int) int {
	return (a * b)

}

// func Eord(b int ) int{
// 	return()
// }
