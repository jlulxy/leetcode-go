package main

import "fmt"

//思路先，排序
func subsetsWithDup(nums []int) [][]int {
	quickSort(nums)
	reData := make([][]int,1)
	pre := -1000
	count := 1
	for _,num := range nums{
		for _,v:=range reData{
			temp :=make([]int,len(v))
			copy(temp,v)
			temp=append(temp,num)
			if pre != num {
				count =1
				reData = append(reData,temp)
			// 在添加是重复的元素只有在前几个重复的数字第一个正常，后面的只有在重复的元素都选了的情况下才能选择
			}else if pre == num && len(v)>=count && v[len(v)-count] == num{
				reData = append(reData,temp)
			}

		}
		if pre == num{
			count++
		}
		pre = num
	}
	return reData
}


func parttion(nums[]int ,left,right,pivotIndex int) int {
	prvotValue := nums[pivotIndex]
	nums[right] ,nums[pivotIndex] = nums[pivotIndex],nums[right]
	storeIndex := left
	for i:=left ; i < right ;i++{
		if nums[i] <= prvotValue{
			nums[i],nums[storeIndex] = nums[storeIndex],nums[i]
			storeIndex++
		}
	}
	nums[right],nums[storeIndex] = nums[storeIndex],nums[right]
	return  storeIndex
}

func quickSort(nums []int){
	if len(nums) <= 1 {
		return
	}
	quickSortre(nums,0,len(nums)-1)
}

func quickSortre(nums[]int,left,right int)  {
	if left < right{
		index:=parttion(nums,left,right,left)
		quickSortre(nums,left,index-1)
		quickSortre(nums,index+1,right)
	}
}

func main(){
	data:=[]int{1,2,2,2}
	re := subsetsWithDup(data)
	fmt.Println(re)
}