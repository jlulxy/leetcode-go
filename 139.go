package main

//Word Break
func setHashMap(wordDict []string)map[string]bool{
	reMap := make(map[string]bool,len(wordDict))
	for i := 0;i < len(wordDict);i++{
		reMap[wordDict[i]] = true
	}
	return reMap
}

//递归时间超限，思想是递归处理每个字符串匹配如果，发现有word匹配记得取
func wordBreakv1(s string, wordDict []string) bool {
	hashMap := setHashMap(wordDict)
	return word_break(s,hashMap,0)
}

func word_break(s string,wordMap map[string]bool,start int) bool{
	if start == len(s){
		return true
	}
	for end:=start+1;end<=len(s);end++{
		stemp:=s[start:end]
		if _, ok := wordMap[stemp]; ok && word_break(s,wordMap,end){
			return  true
		}
	}
	return false
}

//记忆化递归，记录已经计算过的状态减少重复计算
/*
执行用时 :4 ms, 在所有 Go 提交中击败了13.60%的用户
内存消耗 :2.3 MB, 在所有 Go 提交中击败了22.86%
的用户
时间复杂度：O(n^2)O(n
2
 ) 。回溯树的大小最多达到 n^2n
2
  。

空间复杂度：O(n)O(n) 。回溯树的深度可以达到 nn 级别。
 */
func wordBreakMeme(s string, wordDict []string) bool {
	hashMap := setHashMap(wordDict)
	meme := make(map[int]bool)
	return word_break_memp(s,hashMap,0,meme)
}
func word_break_memp(s string,wordMap map[string]bool,start int,meme map[int]bool) bool{
	if start == len(s){
		return true
	}
	if _,ok := meme[start];ok {
		return meme[start]
	}
	for end:=start+1;end<=len(s);end++{
		stemp:=s[start:end]
		if _, ok := wordMap[stemp]; ok && word_break_memp(s,wordMap,end,meme){
			return  true
		}
	}
	meme[start]=false
	return meme[start]
}

/*
* dp[i] = dp[j] && s[j:i]被wordDict匹配 0<j<i
 */
func wordBreak(s string, wordDict []string) bool{
	hashMap := setHashMap(wordDict)
	dp := make([]bool,len(s) + 1)
	dp[0] = true
	for i:=1;i<len(dp);i++{
		dp[i] = false
	}
	for i:=0;i<=len(s);i++{
		for j:=0;j<i;j++{
			stemp :=s[j:i]
			if t,ok := hashMap[stemp];ok && t && dp[j]{
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}



func main (){
	s := "leetcode"
	wordDict := []string{"leet","code"}
	println(wordBreak(s,wordDict))
}