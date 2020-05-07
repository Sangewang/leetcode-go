#!/usr/bin/env python
#-*- coding: utf-8 -*- 

'''
    请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
    输入: "abcabcbb"
    输出: 3 
    解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

    输入: "bbbbb"
    输出: 1
    解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

    输入: "pwwkew"
    输出: 3
    解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
         请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
'''
class Solution(object):
    def __init__(self):
        pass

    def lengthOfLongestSubstring(self, demoStr):
        """
            以demoStr = abcabcbb为例
        """
        #0、求数组长度
        _len = len(demoStr)
        #1、数组长度为0的情况
        if _len == 0:
            return _len
        #2、从头开始计算
        index = 0 #标记开始计算的位置
        
        tmpStrlen = 0 # 记录临时字符串的长度 
        maxStrlen = 0 # 最终最长的字符串
        finalStr = 0
        
        while index < _len:
            hashMap = {} #需要一个hashMap来记录一下映射关系
            tmpStr = '' # 临时记录最长不重复的字串
            _flag = 0
            for i in range(index, _len):
                if demoStr[i] not in tmpStr:
                    tmpStr += demoStr[i]
                    tmpStrlen = len(tmpStr)
                    hashMap[demoStr[i]] = i
                    _flag = 1
                else: #如果发现开始出现重复了
                    print hashMap, demoStr[i]
                    _flag = 0
                    # maxStrlen = max(maxStrlen, tmpStrlen)
                    if tmpStrlen > maxStrlen:
                        finalStr = tmpStr
                        maxStrlen = tmpStrlen
                    index = hashMap[demoStr[i]] + 1
                    break
            # 最后一次循环完了
            if  index == _len-1 or _flag == 1:

                # maxStrlen = max(maxStrlen, tmpStrlen)
                if tmpStrlen > maxStrlen:
                    finalStr = tmpStr
                    maxStrlen = tmpStrlen
                index = hashMap[demoStr[i]] + 1
        return maxStrlen, finalStr


if __name__ == '__main__':
    solve_obj = Solution()
    # demoStr = "abcabcbb"
    # demoStr = ''
    # demoStr = 'bbbbb'
    demoStr = 'pwwkew'
    print solve_obj.lengthOfLongestSubstring(demoStr)