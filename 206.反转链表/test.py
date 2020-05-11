#!/usr/bin/env python
#-*-coding:utf-8-*-

class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def __init__(self):
        pass

    def ReverseList(self, pHead):
        """
            :type head: ListNode
            :rtype: ListNode
            1 -> 2 -> 3 -> 4 -> 5
        """
        pPrev = None
        while pHead:
            pCurr = pHead
            pHead = pHead.next
            pCurr.next = pPrev
            pPrev = pCurr
        return pPrev
