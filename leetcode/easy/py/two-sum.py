class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        whatINeed = {}
        for i in range(1,len(nums)):
            whatINeed[target-nums[i-1]] = i-1
            if nums[i] in whatINeed:
                return [i, whatINeed[nums[i]]]
