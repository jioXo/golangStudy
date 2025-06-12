package main

import "time"

func main() {

	//题目1
	ch := make(chan int)   // 创建一个通道
	go generateNumbers(ch) // 启动生成整数的协程
	printNumbers(ch)       // 主协程接收并打印数据

	//题目2
	bufferedChannelExample() // 调用缓冲通道示例

}

/*
*
题目1 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func generateNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i // 将整数发送到通道中
	}
	close(ch) // 关闭通道，表示不再发送数据
}
func printNumbers(ch <-chan int) {
	for num := range ch { // 从通道中接收数据
		println(num) // 打印接收到的整数
	}
}

/*
*
题目2 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func bufferedChannelExample() {
	ch := make(chan int, 10) // 创建一个缓冲通道，容量为10

	// 生产者协程
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i // 向通道发送整数
		}
		close(ch) // 关闭通道，表示不再发送数据
	}()

	// 消费者协程
	go func() {
		for num := range ch { // 从通道中接收数据
			println(num) // 打印接收到的整数
		}
	}()
	// 等待消费者协程完成
	select {
	case <-time.After(3 * time.Second):
		// 3秒后自动解除阻塞
	} // 阻塞主协程，等待消费者协程完成
}
