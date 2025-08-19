package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
var sw sync.WaitGroup

func write(ch chan int) {
	defer sw.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func read(ch chan int) {
	defer sw.Done()
	for v := range ch {
		fmt.Printf("read:%v\n", v)
	}
}

func main() {
	ch := make(chan int, 10)
	sw.Add(1)
	go write(ch)
	sw.Add(1)
	go read(ch)
	sw.Wait()
}
