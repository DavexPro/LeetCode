class Solution(object):
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        result = []
        i = 1
        while i <= numRows:
            tmp = []
            if i == 1:
                tmp.append(i)
            else:
                j = 0
                while j < i:
                    if j == 0 or j == i - 1:
                        tmp.append(1)
                    else:
                        tmp.append(result[i-2][j-1]+result[i-2][j])
                    j+=1
            result.append(tmp)
            i+=1
        return result
