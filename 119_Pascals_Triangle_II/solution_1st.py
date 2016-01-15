class Solution(object):
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        n = 0
        result = []
        while n <= rowIndex:
            if n == 0:
                result.append(1)
            else:
                tmp = rowIndex + 1 - n
                tmp = result[n-1]*tmp / n
                result.append(tmp)
            n+=1
        return result
