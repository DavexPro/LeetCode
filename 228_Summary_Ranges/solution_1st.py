class Solution(object):
    def summaryRanges(self, nums):
        """
        :type nums: List[int]
        :rtype: List[str]
        """
        n = len(nums)
        i = 0
        result = []
        while i < n:
            if i + 1 == n:
                result.append(str(nums[i]))
                break
            if nums[i] != nums[i + 1] - 1:
                result.append(str(nums[i]))
            else:
                j = 1
                while j < n - i:
                    if nums[i] + j != nums[i+j]:
                        break
                    else:
                        j+=1
                result.append(str(nums[i])+'->'+str(nums[i+j-1]))
                i = i + j - 1
            i+=1
        return result
