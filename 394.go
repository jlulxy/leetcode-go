package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	re,_ :=decodeStringre(s,0)
	return re
}

// 【】一定是匹配的 所以一点遇到']' 可直接跳出递归函数
func decodeStringre(s string,start int ) (string,int){
	restr := ""
	// 结尾标示退出递归
	for j:=start;j<len(s);{
		if s[j] == ']'{
			return restr,j
		}
		if isSChar(rune(s[j])){
			restr += s[j:j+1]
		}
		if isNum(rune(s[j])){
			numStr := s[j:j+1]
			j = j+1
			for isNum(rune(s[j])){
				numStr += s[j:j+1]
				j++
			}
			// 起始标示进入递归
			if num,err :=strconv.Atoi(numStr);err==nil{
				if s[j] == '['{
					tmpStr,end := decodeStringre(s,j + 1)
					for i:=0;i<num;i++{
						restr = restr + tmpStr
					}
					j = end
				}
			}

		}
		j++
	}
	 // 其它条件加入
	 return restr,len(s)
}

func isSChar(s rune)bool{
	if 'A' <= s && s <='z'{
		return true
	}
	return false
}

func isNum(s rune) bool {
	if '0'<= s && '9' >= s{
		return true
	}
	return false
}

func main(){
	fmt.Println(decodeString("20[a]2[bc]"))
}