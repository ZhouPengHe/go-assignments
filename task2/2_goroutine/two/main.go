package main

import (
	"fmt"
	"sync"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
type Task struct {
	Name string
	Job  func()
}

var sw sync.WaitGroup

func main() {
	tasks := []Task{
		{"批量给客户经理发送邮件", func() { time.Sleep(2 * time.Second) }},
		{"批量给客户经理发送飞书", func() { time.Sleep(1 * time.Second) }},
		{"批量给客户经理发送短信", func() { time.Sleep(1 * time.Second) }},
	}

	for _, task := range tasks {
		sw.Add(1)
		go func(t Task) {
			defer sw.Done()
			start := time.Now()
			t.Job()
			end := time.Since(start)
			fmt.Printf("%v:耗时:%v秒\n", task.Name, end.Milliseconds())
		}(task)
	}
	sw.Wait()
}
