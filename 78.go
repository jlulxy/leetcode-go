package main

import "fmt"

//先写暴力的方式

func subsets(nums []int) [][]int{
	return  subSetsRe(nums,3)
}
func subSetsRe (nums[]int,lenSize int)[][]int{
	redata := make([][]int,1)
	if lenSize == 0{
		return redata
	}
	for j:=0;j<len(nums);j++{
		if j+1<=len(nums){
			re:=subSetsRe(nums[j+1:len(nums)],lenSize-1)
			for _,v:=range re{
				v = append(v,nums[j])
				redata = append(redata,v)
			}
		}

	}
	return redata
}

func subsets2(nums[]int) [][]int{
	reData := make([][]int,1)
	for _,num := range nums{
		for _,v:=range reData{
			temp :=make([]int, len(v))
			copy(temp,v)
			temp=append(temp,num)
			reData = append(reData,temp)
		}
	}
	return reData
}
func main(){
	data := []int{0,9,3,5,7}
	fmt.Println(subsets2(data))
	//fmt.Println(data[4:4])
	//fmt.Println(subSetsRe(data,3))
}
