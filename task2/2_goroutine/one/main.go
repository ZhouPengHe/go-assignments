package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
var sg sync.WaitGroup

func printOddNum() {
	defer sg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("奇数：%d\n", i)
			continue
		}
	}
}

func printEven() {
	defer sg.Done()
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("偶数：%d\n", i)
			continue
		}
	}
}

func main() {
	sg.Add(2)
	go printOddNum()
	go printEven()
	sg.Wait()
}
