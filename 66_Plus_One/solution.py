class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        flag = 1
        for i in range(len(digits)-1, -1, -1):
            digits[i] = (digits[i] + 1) % 10 
            if digits[i]:
                flag = 0
                break 
        if flag: 
            digits.insert(0,1)
        return(digits)
