#!/usr/bin/env python
#-*- coding: utf-8 -*- 

"""
    给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
    示例 1：
    输入: "babad"
    输出: "bab"
    注意: "aba" 也是一个有效答案。

    示例 2：
    输入: "cbbd"
    输出: "bb"
"""

class Solution(object):
    def __init__(self):
        pass

    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        # 从头开始遍历，找中心位置进行扩散
        """
        _len = len(s)
        if _len < 2:
            return s
        maxlen = 0
        finalStr = ''
        # 注意这里要分下奇数和偶数
        for i in range(_len):
            # 1、奇数探索
            oddStr = self.findCurrLenPalindrome(s, i, i)
            # 2、偶数探索
            evenStr = self.findCurrLenPalindrome(s, i, i+1)
            # 3、求当前最大
            if len(oddStr) > len(evenStr):
                currlen = len(oddStr) 
                currStr = oddStr
            else:
                currlen = len(evenStr) 
                currStr = evenStr
            if currlen > maxlen:
                maxlen = currlen
                finalStr = currStr
        print finalStr, maxlen
            

    def findCurrLenPalindrome(self, s, left, right):
        i = left
        j = right
        _len = len(s)
        targetStr = ''
        while i >=0 and j < _len:
            if s[i] == s[j]:
                targetStr = s[i:j+1]
                i -= 1
                j += 1
            else:
                break
        return targetStr


if __name__ == '__main__':
    solve_obj = Solution()
    demoStr = "babad"
    demoStr = "cbbd"
    solve_obj.longestPalindrome(demoStr)