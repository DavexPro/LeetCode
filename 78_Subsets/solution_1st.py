class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result = []
        n = len(nums)
        nums.sort()
        while n > 0:
            result.extend(list(itertools.combinations(nums,n)))
            n-=1
        result.append([])
        return result
