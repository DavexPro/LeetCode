class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        flag = 0
        n = len(digits) - 1
        while n >= 0:
            if(digits[n] < 9):
                digits[n] = digits[n] + 1
                break
            else:
                flag = 1
                digits[n] = 0
            n-=1
        if digits[0] == 0:
            digit = [1]
            digit.extend(digits)
            digits = digit
        return digits
