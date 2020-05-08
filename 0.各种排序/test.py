#!/usr/bin/env python
#-*- coding: utf-8 -*- 

import copy
import random

class Solution(object):
    def __init__(self):
        pass

    # [3, 1, 4, 2, 6, 8, 7, 9, 5, 0]
    def quickSort(self, arr, start, end):
        if end <= start:
            return
        left = start
        right = end
        target = arr[left]
        while left < right:
            while left < right and arr[right] > target:
                right -= 1
            arr[left] = arr[right]
            while left < right and arr[left] < target:
                left += 1
            arr[right] = arr[left]
        arr[left] = target
        print arr, left
        self.quickSort(arr, start, left - 1)
        self.quickSort(arr, left + 1, end)
        

    def bubbleSort(self, arr):
        # 冒泡排序
        g_count = 0
        _len = len(arr)
        for i in range(_len - 1):
            for j in range(0, _len - i - 1):
                g_count += 1
                if arr[j] > arr[j+1]:
                    arr[j], arr[j+1] = arr[j+1], arr[j]
        print arr, g_count

    def bubbleSortBest(self, arr):
        # 冒泡排序升级版
        g_count = 0
        _len = len(arr)       
        for i in range(_len - 1):
            _flag = 1
            for j in range(0, _len - i - 1):
                g_count += 1
                if arr[j] > arr[j+1]:
                    arr[j], arr[j+1] = arr[j+1], arr[j]
                    _flag = 0
            if _flag == 1:
                break
        print arr, g_count
            

if __name__ == '__main__':
    # arr = [3, 1, 4, 2, 6, 8, 7, 9, 5, 0]
    # arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    arr = [] 
    for i in range(100):
        arr.append(i)
    random.shuffle(arr)

    solve_obj = Solution()
    solve_obj.bubbleSort(copy.deepcopy(arr))
    solve_obj.bubbleSortBest(copy.deepcopy(arr))

    start = 0
    end = len(arr) - 1
    print "before arr = ", arr
    print solve_obj.quickSort(arr, start, end)
    print "after arr = ", arr

