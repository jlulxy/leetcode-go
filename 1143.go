package main

import "fmt"

//最长公共子序列

func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[i][j] = 1 + d[i-1][j-1]  s1[i] =s2[j]
	//dp[i][j] = max(dp[i-1][j],dp[i][j-1]) s1[i] != s2[j]
	dp := make([][]int,len(text1)+1)
	for i:=0;i<len(text1) +1;i++{
		dp[i] = make([]int,len(text2)+1)
		for j:=0;j<len(text2)+1;j++{
			if i==0 ||j==0{
				dp[i][j] = 0
			}
		}
	}
	for i:=0;i<len(text1);i++{
		for j:=0;j<len(text2);j++{
			if text1[i]==text2[j]{
				dp[i+1][j+1] = dp[i][j]+1
			}else{
				dp[i+1][j+1] = MAX(dp[i][j+1],dp[i+1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func MAX (a,b int) int{
	if a>b {
		return a
	}
	return b
}

func main (){
	text1:="abcde"
	text2:="ace"
	fmt.Println(longestCommonSubsequence(text1,text2))
}