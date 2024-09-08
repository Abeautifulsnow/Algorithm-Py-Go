from typing import List


class Solution:
    def findIntersectionValues(self, nums1: List[int], nums2: List[int]) -> List[int]:
        nums1_set, nums2_set = set(nums1), set(nums2)

        return [
            sum(n1 in nums2_set for n1 in nums1_set),
            sum(n2 in nums1_set for n2 in nums2_set),
        ]
