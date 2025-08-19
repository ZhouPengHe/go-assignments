package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
var sw sync.WaitGroup

type counter struct {
	sm    sync.Mutex
	count int
}

func (c *counter) Inc() {
	c.sm.Lock()
	defer c.sm.Unlock()
	c.count++
}

func (c *counter) GetCount() int {
	c.sm.Lock()
	defer c.sm.Unlock()
	return c.count
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
