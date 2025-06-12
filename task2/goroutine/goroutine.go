package main

import (
	"sync"
	"time"
)

func main() {

	//题目1
	var wg sync.WaitGroup
	wg.Add(2) // 等待两个协程
	go func() {
		printOdd()
		wg.Done()
	}()
	go func() {
		printEven()
		wg.Done()
	}()
	wg.Wait() // 等待两个协程结束

	//题目2
	tasks := []func(){
		func() { time.Sleep(1 * time.Second); println("Task 1 completed") },
		func() { time.Sleep(2 * time.Second); println("Task 2 completed") },
		func() { time.Sleep(3 * time.Second); println("Task 3 completed") },
	}
	taskScheduler(tasks) // 调度并执行任务
}

/**
题目1 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/

func printOdd() {
	for i := 1; i <= 10; i += 2 {
		println(i) // 打印奇数
	}
}
func printEven() {
	for i := 2; i <= 10; i += 2 {
		println(i) // 打印偶数
	}
}

/*
*
题目2 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func taskScheduler(tasks []func()) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)
			println("Task executed in:", duration.String())
		}(task)
	}
	wg.Wait() // 等待所有任务完成
}
