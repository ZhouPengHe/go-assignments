package main

import "fmt"

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。

func slicesMulTwo(slice *[]int) {
	for index, val := range *slice {
		(*slice)[index] = val * 2
	}
}

func main() {
	slice := []int{1, 2, 3}
	slicesMulTwo(&slice)
	fmt.Println(slice)
}
