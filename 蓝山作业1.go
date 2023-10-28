package main

import (
	"fmt"
	"math"
)

func main() {

	var r float64
	fmt.Print("请输入圆的半径：r = ")
	fmt.Scanf("%f", &r)
	S := math.Pow(r, 2) * math.Pi
	fmt.Print("圆的面积为：S = ", S)
}
