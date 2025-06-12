package main

import "fmt"

func main() {
	var strs []string = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	//长度为1或者0都直接返回
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	//从第二个字符串开始，和第一个字符串比较
	//如果第一个字符串的某个位置和第二个字符串不同，就把第一个字符串截断到这个位置
	for i, v := range strs {
		if i == 0 {
			continue
		}
		for j := 0; j < len(strs[0]) && j < len(v); j++ {
			if strs[0][j] != v[j] {
				strs[0] = strs[0][:j]
				break
			}
		}

	}
	return strs[0]
}
