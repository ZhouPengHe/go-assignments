package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
var sw sync.WaitGroup

type counter struct {
	count atomic.Int64
}

func (c *counter) Inc() {
	c.count.Add(1)
}

func (c *counter) GetCount() int64 {
	return c.count.Load()
}

func main() {
	goroutineCount := 10
	c := counter{}
	sw.Add(goroutineCount)
	for i := 1; i <= goroutineCount; i++ {
		go func() {
			defer sw.Done()
			for j := 0; j < 1000; j++ {
				c.Inc()
			}
		}()
	}
	sw.Wait()
	fmt.Println(c.GetCount())
}
