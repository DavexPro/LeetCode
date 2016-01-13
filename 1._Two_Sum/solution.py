class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        # sort
        sorted_num = sorted(nums)

        # two points
        left = 0
        right = len(nums) - 1
        res = []
        while (left < right):
            sum = sorted_num[left] + sorted_num[right]
            if sum == target:
                # find out index
                break;
            elif sum > target:
                right -= 1
            else:
                left += 1

        if left == right:
            return -1, -1
        else:
            pos1 = nums.index(sorted_num[left]) + 1
            pos2 = nums.index(sorted_num[right]) + 1
            if pos1 == pos2:    # find again
                pos2 = nums[pos1:].index(sorted_num[right]) + pos1 + 1
            return min(pos1, pos2), max(pos1, pos2)
