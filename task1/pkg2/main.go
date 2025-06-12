package main

import "fmt"

func main() {
	ints := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(ints))
	//fmt.Println(myMethod(ints))
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//思路，如果遍历值相等则删除数组中的这个元素
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	//返回新数组的长度
	fmt.Println(nums)
	nums = nums[:i+1]
	fmt.Println(nums)
	return len(nums)
}

func myMethod(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用 map 记录已存在的元素（优化查找效率）
	seen := make(map[int]bool)
	newints := []int{}

	for _, v := range nums {
		// 如果元素不存在于 map 中，则添加到结果切片
		if !seen[v] {
			seen[v] = true
			newints = append(newints, v)
		}
	}
	fmt.Println(newints)
	fmt.Println(nums)
	return len(newints)
}
