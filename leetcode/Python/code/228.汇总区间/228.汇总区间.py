from typing import List


class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        if len(nums) <= 1:
            return list(map(str, nums))

        res = []
        start: int = 0

        for i, el in enumerate(nums):
            if i == 0:
                if el + 1 == nums[i + 1]:
                    start = el
                else:
                    res.append(str(el))
            elif i > 0 and i < len(nums) - 1:
                if el - 1 != nums[i - 1] and el + 1 != nums[i + 1]:
                    res.append(str(el))
                else:
                    if el - 1 == nums[i - 1] and el + 1 != nums[i + 1]:
                        res.append(f"{start}->{el}")
                    if el - 1 != nums[i - 1] and el + 1 == nums[i + 1]:
                        start = el
            else:
                if el - 1 == nums[i - 1]:
                    res.append(f"{start}->{el}")
                else:
                    res.append(str(el))

        return res

    def summaryRangesM1(self, nums: List[int]) -> List[str]:
        l, r = 0, 1  # noqa: E741
        res = []

        while l < len(nums):
            while r < len(nums) and nums[r] == nums[r - 1] + 1:
                r += 1
            if l == r - 1:
                res.append(str(nums[l]))
            else:
                res.append(str(nums[l]) + "->" + str(nums[r - 1]))
            # 注意这里在一个区间处理完毕之后，右端点r还是要+1的
            l = r  # noqa: E741
            r += 1

        return res


s = Solution()
print(s.summaryRangesM1([0, 1, 2, 4, 5, 7]))
