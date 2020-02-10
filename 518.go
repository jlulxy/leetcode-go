package main

import "fmt"

/**
	完全背包
	dp[i] +=dp[i-cost]
	dp[0] = 1
 */
func change(amount int, coins []int) int {
	dp:=make([]int,amount+1)
	dp[0] = 1
	for _,cost := range coins{
		for i:=0;i<=amount;i++{
			if i>=cost{
				dp[i] += dp[i-cost]
			}
		}
	}
	return dp[amount]
}

func main(){
	data:=[]int{1,2,5}
	fmt.Println(change(5,data))

}