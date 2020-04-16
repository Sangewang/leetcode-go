/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：
输入: "cbbd"
输出: "bb"
*/

package main

import (
	"fmt"
	"leetcode/common"
)

func findLongestCommonStr(src string, dst string) ([]string, int) {
	// 1、构建一个二维数组
	srcLen := len(src) // 作为纵列
	dstLen := len(dst) // 作为横列
	fmt.Println("srcLen = ", srcLen, "dstLen = ", dstLen)
	//var dpArray [srcLen][5]int
	
	// 2、先构建一个行
	dpArray := make([][]int, srcLen + 1)
	
	// 3、每一个行的索引 => 再构建一个列
	for i := range dpArray {
		dpArray[i] = make([]int, dstLen + 1)
	}

	// 4、输出二维数组
	fmt.Println("dpArray = ", dpArray)
	/*
		  b a b a d
		0 0 0 0 0 0 dst
	  d	0 0 0 0 0 1
	  a	0 0 1 0 1 0
	  b	0 1 0 2 0 0
	  a	0 0 2 0 3 0
	  b 0 1 0 3 0 0
	src
	if src[i] == dst[j] => dpArray[i-1][j-1] + 1
	if src[i] != dst[j] => 0
	*/
	// 5、遍历找最大的
	var maxLen int = 0
	for i := 1; i < srcLen + 1; i ++ {
		for j := 1; j < dstLen + 1; j ++ {
			if src[i - 1] == dst[j - 1] {
				dpArray[i][j] = dpArray[i-1][j-1] + 1
				if dpArray[i][j] > maxLen {
					maxLen = dpArray[i][j] 
				}
			} else {
				dpArray[i][j] = 0
			}
		}
	}
	fmt.Println("dpArray = ", dpArray, "maxLen = ", maxLen)

	// 6、查找最大的有哪些
	retStr := ""
	ret := make([]string, 0)
	for i := srcLen; i >= 0; i -- {
		for j := dstLen; j >=0; j -- {
			retStr = ""
			if dpArray[i][j] == maxLen {
				tempI := i
				tempJ := j
				for dpArray[tempI][tempJ] > 0{
					retStr += string(src[tempI-1])
					tempI --
					tempJ --
				}
				retStr = common.Reverse(retStr)
				ret = append(ret, retStr)
				fmt.Println("retStr = ", retStr)
			}
		}
	}
	return ret, maxLen
}

// 最长的回文
func longestPalindrome(s string) string {
	// 1、先做翻转
	fmt.Println("标准的回文输入 = ", s)
	reverS := common.Reverse(s)
	fmt.Println("标准的回文翻转 = ", reverS)
	// 2、求最长连续的公共子序列
	ret, maxLength := findLongestCommonStr(s, reverS)
	fmt.Println("longestPalindrome maxLength = ", maxLength, "ret = ", ret)
	return string(ret[0])
}

func main () {
	fmt.Println("练习回文算法")
	var s string
	// s = "babad"
	// s = "cbbd"
	// s = "abb"
	s = "aacdefcaa"
	ret := longestPalindrome(s)
	fmt.Println("返回结果 = ", ret)
}

