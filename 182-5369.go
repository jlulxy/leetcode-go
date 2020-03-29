package main

import "fmt"

func numTeams(rating []int) int {
	reData := 0
	for i :=0 ;i < len(rating);i++{
		for j :=i+1;j<len(rating);j++{
			if rating[i] <=rating[j]{
				continue
			}
			for k:=j+1;k<len(rating);k++{
				if rating[j] <= rating[k]{
					continue
				}
				reData++
			}
		}
	}

	for i :=0 ;i < len(rating);i++{
		for j :=i+1;j<len(rating);j++{
			if rating[i] >=rating[j]{
				continue
			}
			for k:=j+1;k<len(rating);k++{
				if rating[j] >= rating[k]{
					continue
				}
				reData++
			}
		}
	}
	return reData
}
func main(){
	data := []int{1,2,3,4}
	fmt.Println(numTeams(data))
}
