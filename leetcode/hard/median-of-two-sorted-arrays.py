class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        for _ in nums2:
            nums1.append(_)
        nums1.sort()
        if len(nums1)%2:
            return nums1[len(nums1)//2]
        else:
            return (nums1[len(nums1)//2-1]+nums1[len(nums1)//2])/2
            
