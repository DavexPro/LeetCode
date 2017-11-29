class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        
        if len(nums) == 0:
            return None
        
        nums_dict = {}
        
        for i, value in enumerate(nums):
            nums_dict[value] = i
        
        for i in range(len(nums)):
            last = target - nums[i]
            if last in nums_dict and nums_dict[last] != i:
                return [i, nums_dict[last]]
        
        return None
        
        
