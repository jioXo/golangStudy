package main

import "fmt"

func main() {

	ints := []int{1, 2, 3,9,9}
	result := plusOne(ints)
	fmt.Println(result)
}

/*
*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    // 如果所有位都是9，最后需要在最前面加1
    return append([]int{1}, digits...)
}
