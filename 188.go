package main

import "fmt"

//要注意 K次交易代表的是最多K次交易所以做basecase的时候要考虑好

func maxProfit188(k int, prices []int) int {
	//转移方程dp[i][j][k]  i 天，j交易次数()，k持有状态
	//dp[i][j][0] = max(dp[i-1][j][0],dp[i-1][j-1][1] + prices[i]
	//dp[i][j][1] = max(dp[i-1][j][1],dp[i-1][j][0] - prices[i])
	if len(prices)==0 {
		return 0
	}
	if k> len(prices)/2{
		return maxProfileFG(prices)
	}
	dp := make([][][]int,len(prices))

	for i :=0 ;i<len(prices);i++{
		dp[i] = make([][]int,k+1)
		for j := 0;j <= k; j++{
			dp[i][j] = make([]int,2)
			for h :=0;h<2;h++{
				if j==0 &&h==1{
					dp[i][j][h] = -10000
				}else if i==0&& h==1{
					dp[i][j][h] = -prices[i]
				}
			}
		}
	}
	/*
	dp[0][0][0] = 0
	dp[0][0][1] = -100000
	dp[0][1][0] = 0
	dp[0][1][1] = -prices[0]
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]
	*/
	for i:=1;i<len(prices);i++{
		for j:=k ;j>=1; j--{
			dp[i][j][0] = max188(dp[i-1][j][0],dp[i-1][j][1] + prices[i])
			dp[i][j][1] = max188(dp[i-1][j][1],dp[i-1][j-1][0] - prices[i])
		}
	}

	return dp[len(prices)-1][k][0]
}
func maxProfit188Spcae (k int, prices []int) int{
	//转移方程dp[i][j][k]  i 天，j交易次数()，k持有状态
	//dp[i][j][0] = max(dp[i-1][j][0],dp[i-1][j-1][1] + prices[i]
	//dp[i][j][1] = max(dp[i-1][j][1],dp[i-1][j][0] - prices[i])
	if len(prices)==0 {
		return 0
	}
	if k> len(prices)/2{
		return maxProfileFG(prices)
	}
	dp := make([][2]int,k+1)
	//由于i是从0开始所以初始化时候是-1天所以-1天是不能有操作和持有的！！！！！！
	for i:=1;i<k+1;i++{
		dp[i][1]=-10000
	}
	for i:=0;i<len(prices);i++{
		for j:=k ;j>=1; j--{
			dp[j][1] = max188(dp[j][1],dp[j-1][0] - prices[i])
			dp[j][0] = max188(dp[j][0],dp[j][1] + prices[i])
		}
	}

	return dp[k][0]

}

func maxProfileFG(prices[]int) int {
	if len(prices) == 0{
		return 0
	}
	maxprofit , valley,peek, i := 0,prices[0],prices[0] ,0
	for i<len(prices)-1{
		for i< len(prices)-1 && prices[i] >=prices[i+1]{
			i++
		}
		valley = prices[i]
		for i< len(prices)-1 && prices[i] <= prices[i+1]{
			i++
		}
		peek = prices[i]
		maxprofit += peek -valley
	}
	return  maxprofit
}

func max188(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main(){
	data := []int{2,1,4}
	fmt.Println(maxProfit188Spcae(2,data))
}