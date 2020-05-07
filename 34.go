package main

import "fmt"

func searchRangeV2(nums []int, target int) []int {

	l, r := 0, len(nums)-1
	if r < 0 {
		return []int{-1, -1}
	}
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	res1 := l
	l, r = 0, len(nums)-1
	for l < r {
		mid := (l+r)/2 + 1
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid
		}
		fmt.Println(l, r)
	}
	res2 := l
	if nums[res1] != target {
		return []int{-1, -1}
	} else {
		return []int{res1, res2}
	}
}

func searchRange(nums []int, target int) []int {
	left := binaryLeftSerach(nums, target)
	right := binaryRightSerach(nums, target)
	return []int{left, right}
}

func binaryLeftSerach(nums []int, target int) int {
	lenN := len(nums)
	if lenN == 0 {
		return -1
	}
	right, left, mid := lenN, 0, 0
	for left < right {
		mid = (left + right) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left == lenN {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func binaryRightSerach(nums []int, target int) int {
	lenN := len(nums)
	if lenN == 0 {
		return -1
	}
	right, left, mid := lenN, 0, 0
	for left < right {
		mid = (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if right-1 < 0 {
		return -1
	}
	if nums[right-1] == target {
		return right - 1
	}
	return -1
}

func main() {
	data := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(data, 8))
}
