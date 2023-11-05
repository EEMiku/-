package main

import "fmt"

type Operation func(int, int) int

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	return a / b
}
func Calculator(a, b int, op Operation) int {
	return op(a, b)
}
func main() {
	var a, b int
	fmt.Print("请输入两个数(中间用空格隔开)；")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d + %d = %d\n", a, b, Calculator(a, b, Add))
	fmt.Printf("%d - %d = %d\n", a, b, Calculator(a, b, Sub))
	fmt.Printf("%d * %d = %d\n", a, b, Calculator(a, b, Mul))
	fmt.Printf("%d / %d = %d\n", a, b, Calculator(a, b, Div))
}
