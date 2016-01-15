class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums) - 1
        i = 0
        while i < n:
            n = len(nums) - 1
            if i + 1 <= n:
                if nums[i] == nums[i+1]:
                    del nums[i]
                else:
                    i+=1
