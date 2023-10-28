package main

import (
	"fmt"
	"math"
)

func Circlearea(r float64) float64 {
	S := math.Pow(r, 2) * math.Pi
	return S
}

func main() {

	var r float64
	fmt.Print("请输入圆的半径：r = ")
	fmt.Scanf("%f", &r)
	S := Circlearea(r)
	fmt.Print("圆的面积为：S = ", S)
}
