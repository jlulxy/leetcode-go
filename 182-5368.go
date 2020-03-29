package main

import "fmt"

func findLucky(arr []int) int {
	checkMap := make(map[int]int)
	for _,key:= range arr{
		checkMap[key] +=1
	}
	ret := -1
	for k,v := range checkMap{
		if k == v && k>=ret {
			ret = k

		}
	}
	return ret
}

func main (){
	data := []int{7,7,7,7,7,7,7}
	fmt.Println(findLucky(data))
}