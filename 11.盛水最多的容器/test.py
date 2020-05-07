#!/usr/bin/env python
#-*- coding: utf-8 -*- 
"""
Area = Max(min(height[i], height[j]) * (j-i)) {其中0 <= i < j < height,size()}

输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
8,6,2,5,4,8,3,7 => 49
"""
class Solution(object):
    def __init__(self):
        pass

    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        # 双指针 求面积
        _len = len(height)
        if _len == 0:
            return 0
        left = 0
        maxArea = 0
        right = _len - 1
        while left < right:
            h = min(height[left], height[right])
            curArea = h * (right - left)
            maxArea = max(curArea, maxArea)
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return maxArea

if __name__ == '__main__':
    sovle_obj = Solution()
    arr = [1,8,6,2,5,4,8,3,7]
    print sovle_obj.maxArea(arr)