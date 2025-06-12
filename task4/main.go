package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

/**
两数之和
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nums
}
