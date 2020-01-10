package main

import "fmt"


func maximalSquare(matrix [][]byte) int {
	//使用dp转移
	//dp[i][j] = min(dp[i-1,j-1],dp[i-1],[j],dp[j]) +1
	rows := len(matrix)
	if rows <= 0 {
		return 0
	}
	cols:= len(matrix[0])
	dp:= make([][]int,len(matrix))
	len := 0
	//初始化dp数组
	for i:=0;i<rows;i++{
		dp[i] = make([]int,cols)
		for j :=0 ;j<cols;j++{
			if matrix[i][j] == '1' {
				if i==0||j==0 {
					dp[i][j] = 1
				}else{
					dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1]) +1
				}
				}else{
				dp[i][j]=0
			}
			if dp[i][j] > len{
				len = dp[i][j]
			}
		}
	}
	return len *len
}
func min(a,b,c int) int {
	if a < b{
		b  = a
	}
	if b < c {
		return  b
	}else{
		return c
	}
}
func main()  {
	matrix := make([][]byte,3)
	for i:=0;i<3;i++{
		matrix[i] = make([]byte,3)
		for j:=0;j<3;j++{
			matrix[i][j] = '0'+1
		}
	}
	fmt.Println(matrix)
	//ret :=maximalSquare(matrix)
	//fmt.Println(ret)
}
