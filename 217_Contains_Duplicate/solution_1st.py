class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        nums = sorted(nums)
        n = len(nums)
        i = 0
        while i < n:
            if i + 1 == n:
                return False
            if nums[i] == nums[i + 1]:
                return True
            else:
                i+=1
        return False
