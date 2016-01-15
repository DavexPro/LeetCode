
class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums_sorted = sorted(nums)
        n = len(nums)/2
        return nums_sorted[n]
