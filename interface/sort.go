package main

import "fmt"

// 排序接口：只要实现了这个三个方法的对象，就可以进行排序
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 进行排序的函数
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

type IntArray []int // 定义新的类型

// 类型实现排序接口需要的三个方法
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	data := []int{1, 4, 6, 2} // 初始化整数数组
	Sort(IntArray(data))      // 进行类型转换和排序
	fmt.Println(data)
}
