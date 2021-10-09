package main

import "fmt"

// 形状接口，包含面积计算方法
type Shaper interface {
	Area() float32
	Perimeter() float32
}

// 正方形结构体
type Square struct {
	side float32
}

// 正方形实现面积方法
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// 正方形实现周长方法
func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

func main() {
	sq1 := new(Square) // 实例一个正方形
	sq1.side = 5

	areaIntf := Shaper(sq1) // 绑定实例与接口，得到接口变量，如果实例方法定义跟接口不同会报错

	fmt.Printf("The square has area: %f\n", areaIntf.Area()) // 只能调用接口拥有的方法，不能操作实例
	fmt.Printf("The square has perimeter: %f\n", areaIntf.Perimeter())

	// 校验正方形实例是否实现了形状接口
	if _, ok := areaIntf.(Shaper); ok {
		fmt.Printf("yes") // note: sv, not v
	}
}
