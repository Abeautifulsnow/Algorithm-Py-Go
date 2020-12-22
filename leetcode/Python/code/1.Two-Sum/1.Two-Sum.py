"""
No.1
题目：两数之和
难度：简单

给定一个整数数组``nums``和一个目标值``target``，请你在该数组中找出和为目标值的那``两个``整数，
并返回他们的数组下标。你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中
同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
"""
from typing import List


class Solution(object):
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        h = {}
        for i, j in enumerate(nums):
            k = target - j
            if k in h:
                return [h[k], i]
            h[j] = i


class Solution2(object):
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i, v in enumerate(nums):
            for j in range(i + 1, len(nums)):
                if v + nums[j] == target:
                    return [i, j]


nums = [5, 5, 5]
target = 10
d = Solution2()
e = d.twoSum(nums, target)
print(*list(e))
