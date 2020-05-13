#!/usr/bin/env python
#-*-coding:utf-8-*-

class Solution(object):
    def __init__(self):
        pass
    
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        1
        1 1
        1 2 1
        1 3 3 1
        1 4 6 4 1
        """

        # ret = [[0 for i in range(numRows)] for j in range(numRows)]
        ret = [[] for i in range(numRows)]
        print ret
        for i in range(numRows):
            # print i 
            ret[i].extend([0 for j in range(i+1)])
        print ret
        
        for i in range(numRows):
            ret[i][0] = 1
            ret[i][i] = 1
        print ret
                
        if numRows <= 2:
            return ret

        for i in range(2, numRows):
            for j in range(1, len(ret[i]) - 1):
                ret[i][j] = ret[i - 1][j - 1] + ret[i - 1][j]
        print ret
        print ret[numRows - 1]
        return ret

        
    def bestGenerate(self, rowIndex):
        res = [0]*(rowIndex+1)
        res[0] = 1
        for i in range(rowIndex+1):
            res[i] = 1
            for j in range(i-1,0,-1):
                res[j] += res[j-1]
        print res
        return res


if __name__ == '__main__':
    solve_obj = Solution()
    solve_obj.bestGenerate(3+1)