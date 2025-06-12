package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//methond1() // 使用互斥锁保护共享计数器
	methond2() // 使用原子操作实现无锁计数器
}

/*
*
题目1 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

func methond1() {
	counter := 0         // 共享计数器
	var mutex sync.Mutex // 创建一个互斥锁
	var wg sync.WaitGroup

	wg.Add(10) // 设置等待10个协程
	for i := 0; i < 10; i++ {
		go func() {
			incrementCounter(&counter, &mutex)
			wg.Done() // 协程完成后计数减一
		}()
	}
	wg.Wait()            // 等待所有协程完成
	fmt.Println(counter) // 输出最终计数器的值
}
func incrementCounter(counter *int, mutex *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()            // 加锁
		*counter = *counter + 1 // 递增
		mutex.Unlock()          // 解锁
	}
}

/*
*
题目2 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func methond2() {
	var counter int64 // 无锁计数器
	var wg sync.WaitGroup

	wg.Add(10) // 设置等待10个协程
	for i := 0; i < 10; i++ {
		go func() {
			incrementCounterAtomic(&counter)
			wg.Done() // 协程完成后计数减一
		}()
	}
	wg.Wait()            // 等待所有协程完成
	fmt.Println(counter) // 输出最终计数器的值
}

func incrementCounterAtomic(counter *int64) {
	for i := 0; i < 1000; i++ {
		// 使用原子操作递增计数器
		atomic.AddInt64(counter, 1)
	}
}
