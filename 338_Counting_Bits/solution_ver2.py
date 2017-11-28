class Solution():
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        results = [0 for i in range(num + 1)]
        
        for i in range(num + 1):
            results[i] = results[i >> 1] + (i & 1)
            
        return results
        
