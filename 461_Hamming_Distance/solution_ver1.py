class Solution:
    def hammingDistance(self, x, y):
        """
        :type x: int
        :type y: int
        :rtype: int
        """
        dis = 0
        while x > 0 or y > 0:
            if x & 1 != y & 1:
                dis += 1
            x = x >> 1
            y = y >> 1
        
        return dis
