package main

import (
	"fmt"
	"math"
)

/*
*
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	r := Rectangle{3, 4}
	fmt.Printf("矩形的面积为：%v，周长为：%v\n", r.Area(), r.Perimeter())

	c := Circle{3.5}
	fmt.Printf("圆的面积为：%.2f，周长为：%.2f\n", c.Area(), c.Perimeter())
}
