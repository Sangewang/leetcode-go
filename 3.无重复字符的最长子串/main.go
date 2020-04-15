/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
	示例 1:
	输入: "abcabcbb"
	输出: 3 
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
	示例 2:
	输入: "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
	示例 3:
	输入: "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
package main 

import (
	"fmt"
	"strings"
)

// 求最长的不重复字符的子串
func lengthOfLongestSubstring(s string) (int, string) {
	strLen := len(s)
	if strLen == 0 {
		return strLen, ""
	}
	fmt.Println("lengthOfLongestSubstring len = ", strLen, " randomStr = ", s)
	commStr := string(s[0])
	maxLen := 1
	currLen := 1
	new_commStr := ""
	finalStr := commStr
	for j := 1 ; j < strLen; j++ {
		// 判断当前的字符是否在字符串中，且位置是哪里
		currStr := string(s[j])
		currIndex := strings.Index(commStr, currStr)
		fmt.Println(currStr, "在", commStr, "的位置", currIndex, " currLen= ", currLen)
		if currIndex == -1 {
			commStr += currStr
			currLen += 1
			if currLen > maxLen {
				finalStr = commStr
				maxLen = currLen
			}
		}else {
			new_commStr = commStr[currIndex + 1:] + currStr
			commStr = new_commStr
			currLen = len(commStr)
		}
		fmt.Println("最终替换之后的commStr = ", commStr, " currLen = ", currLen, "currIndex = ", currIndex, " j = ", j)
	}
	return maxLen, finalStr
}

func main () {
	
	// randomStr := "abcabcbb"
	// randomStr := "bbbb"
	// randomStr := "pwwkew"
	randomStr := "bbtablud"
	maxLength, finalStr := lengthOfLongestSubstring(randomStr)
	fmt.Println("无重复字符的最长子串 maxLength = ", maxLength, " finalStr = ", finalStr)
}