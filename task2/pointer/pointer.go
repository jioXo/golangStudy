package main

import "fmt"

func main() {
	//题目一
	var value int = 5
	incrementValue(&value)            // 传递指针
	println("Modified value:", value) // 输出修改后的值

	//题目二
	var slice []int = []int{1, 2, 3, 4, 5}
	doubleSliceValues(&slice) // 传递切片指针
}

/**
题目1 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/

func incrementValue(ptr *int) {
	*ptr += 10 // 将指针指向的值增加10
}

/*
*
题目2 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func doubleSliceValues(slice *[]int) {
	for i := range *slice {
		fmt.Println("Before doubling:", (*slice)[i]) // 输出修改前的值
		(*slice)[i] *= 2                             // 将切片中的每个元素乘以2
		fmt.Println("After doubling:", (*slice)[i])  // 输出修改后的值
	}
}
