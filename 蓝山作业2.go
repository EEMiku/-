package main

import "fmt"

func Calculator(num1 int, num2 int, CMD func(int, int) int) int {
	return CMD(num1, num2)
}
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

func main() {
	var a, b int
	var c string
	fmt.Print("请输入两个数和运算符(+ - * /)(中间用空格隔开)；")
	fmt.Scanf("%d %d %s", &a, &b, &c)
	switch c {
	case "+":
		fmt.Printf("%d + %d = %d\n", a, b, Calculator(a, b, Add))
	case "-":
		fmt.Printf("%d - %d = %d\n", a, b, Calculator(a, b, Sub))
	case "*":
		fmt.Printf("%d * %d = %d\n", a, b, Calculator(a, b, Mul))
	case "/":
		fmt.Printf("%d / %d = %d\n", a, b, Calculator(a, b, Div))
	}
}
