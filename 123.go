package main

import "fmt"

//要注意 K次交易代表的是最多K次交易所以做basecase的时候要考虑好
func maxProfit123(prices []int) int {
	//转移方程dp[i][j][k]  i 天，j交易次数()，k持有状态
	//dp[i][j][0] = max(dp[i-1][j][0],dp[i-1][j-1][1] + prices[i]
	//dp[i][j][1] = max(dp[i-1][j][1],dp[i-1][j][0] - prices[i])
	if len(prices)==0 {
		return 0
	}
	dp := make([][][]int,len(prices))

	for i :=0 ;i<len(prices);i++{
		dp[i] = make([][]int,3)
		for j := 0;j < 3; j++{
			dp[i][j] = make([]int,2)
		}
	}
	dp[0][0][0] = 0
	dp[0][0][1] = -100000
	dp[0][1][0] = 0
	dp[0][1][1] = -prices[0]
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]
	for i:=1;i<len(prices);i++{
		for j:=2 ;j>=1; j--{
			dp[i][j][0] = max123(dp[i-1][j][0],dp[i-1][j][1] + prices[i])
			dp[i][j][1] = max123(dp[i-1][j][1],dp[i-1][j-1][0] - prices[i])
		}
	}
	if dp[len(prices)-1][2][0] <= 0{
		return 0
	}
	return dp[len(prices)-1][2][0]
}

//空间优化+时间优化
/***
	time o(n)  space 0(1)
	99.49      80.00
 */
func maxProfit123Space(prices []int) int {
	//转移方程dp[i][j][k]  i 天，j交易次数()，k持有状态
	//dp[i][j][0] = max(dp[i-1][j][0],dp[i-1][j-1][1] + prices[i]
	//dp[i][j][1] = max(dp[i-1][j][1],dp[i-1][j][0] - prices[i])
	if len(prices) == 0 {
		return 0
	}
	dp_i_0_0 := 0
	dp_i_0_1 := -10000
	dp_i_1_0 := 0
	dp_i_1_1 := -prices[0]
	for i := 0; i < len(prices); i++ {
		dp_i_1_0 = max123(dp_i_1_0, dp_i_1_1+prices[i])
		dp_i_1_1 = max123(dp_i_1_1, dp_i_0_0-prices[i])
		dp_i_0_0 = max123(dp_i_0_0, dp_i_0_1+prices[i])
		dp_i_0_1 = max123(dp_i_0_1, -prices[i])
	}
	return dp_i_1_0
}


func max123(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main(){
	data := []int{1,2,4,2,5,7,2,4,9,0}
	fmt.Println(maxProfit123Space(data))
}