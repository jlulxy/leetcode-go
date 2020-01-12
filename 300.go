package main

import "fmt"

func lengthOfLIS(nums []int) int {
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		tempLen := 1
		pre := nums[i]
		for j := i; j < len(nums); j++ {
			if nums[j] > pre {
				tempLen++
			}
			pre = nums[j]
		}
		if tempLen > maxLen {
			maxLen = tempLen
		}
	}
	return maxLen
}

func main() {
	input := []int{4,10,4,3,8,9}
	fmt.Println(lengthOfLIS(input))
}
