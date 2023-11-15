package main

import (
	"fmt"
)

type Goods struct {
	Name   string
	Price  float64
	Number int
}

type Manager interface {
	Check() int
	Update(int)
	PrintInformation1() string
}

type Electronics struct {
	Brand string
	Model string
	Goods
}

func (e *Electronics) Check() int {
	return e.Number
}

func (e *Electronics) Update(NewGoods int) {
	e.Number = NewGoods
}

func (e *Electronics) PrintInformation1() {
	fmt.Println("商品名称：%s,库存数量：%d,价格：%.2f", e.Name, e.Number, e.Price)
}

func (e *Electronics) PrintfInformation2() {
	fmt.Println("品牌：%s，型号：%s", e.Brand, e.Model)
}
