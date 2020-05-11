#!/usr/bin/env python
#-*-coding:utf-8-*-

'''
    输入: [3,4,5,1,2]
    输出: 1

    输入: [4,5,6,7,0,1,2]
    输出: 0

'''
class Solution(object):
    def __init__(self):
        pass

    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

    #使用快排可以搞一下
    def quickSort(self, nums, start, end):
        if start > end:
            return
        low = start
        high = end
        target = nums[start]
        while low < high:
            while low < high and nums[high] > target:
                high -= 1
            nums[low] = nums[high]
            while low < high and nums[low] < target:
                low += 1
            nums[high] = nums[low]
        nums[low] = target
        self.quickSort(nums, start, low - 1)
        self.quickSort(nums, low + 1, end)
        return nums[0]
        


if __name__ == '__main__':
    solve_obj = Solution()
    nums = [4,5,6,7,0,1,2]
    # nums = [1]
    print solve_obj.quickSort(nums, 0, len(nums) -1)