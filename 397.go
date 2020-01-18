package main

import (
	"fmt"
)

func integerReplacementV1(n int) int {
	//dp[i] =  min（dp[n-1]+1，dp[n+1]+1）   i =奇数
	//dp[i] = dp[n/2] +1
	//需要注意奇数的时候可能会用到n+1的情形
	if n == 0 {
		return  -1
	}
	if n==1{
		return 0
	}
	if n <=2{
		return 1
	}
	dp := make(map[int]int,n/2)
	for i := 2 ;i <= n+1 ;i++{
		if i  & 0x1 == 0 {
			dp[i] = dp[i/2] + 1
			dp[i-1] = Min(dp[i]+1, dp[i-2]+1)
		}
	}
	return dp[n]

}

func integerReplacement(n int) int {
	count:=0
	for n != 1{
		if n & 0x01 == 0{
			n = n>>1
			count++

		}else{
			if n & 0x02 == 0{
				n=n-1
				count++
			}else{
				if n==3 {
					count=count+2
					return count
				}else{
					n = n+1
					count++
				}
			}
		}

	}
	return  count
}

func Min (a,b int) int {
	if a<b {
		return a
	}
	return b
}

func main(){
	fmt.Println(integerReplacementV1(10000000))
	fmt.Println(integerReplacement(10000000))
}