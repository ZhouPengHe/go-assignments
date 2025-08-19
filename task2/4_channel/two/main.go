package main

import (
	"fmt"
	"sync"
)

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
var sw sync.WaitGroup

func producer(ch chan int) {
	defer sw.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	defer sw.Done()
	for v := range ch {
		fmt.Printf("consumer:%v\n", v)
	}
}

func main() {
	ch := make(chan int, 10)
	sw.Add(1)
	go producer(ch)
	sw.Add(1)
	go consumer(ch)
	sw.Wait()
}
