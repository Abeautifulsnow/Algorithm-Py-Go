from typing import Generator, List


class Solution:
    def containsNearbyAlmostDuplicate(
        self, nums: List[int], indexDiff: int, valueDiff: int
    ) -> bool:
        if len(nums) <= 1 or indexDiff < 0 or valueDiff < 0:
            return False

        # Hash, 处理每个桶
        all_buckets = {}

        # 每个桶的容量
        bucket_size = valueDiff + 1

        for i, v in enumerate(nums):
            # 1. 取商, 决定放入哪个桶
            bucket_num = v // bucket_size

            # 2. 桶中已经有元素了: 说明记录相同桶的上一个值 和 本次的值 都在valueDiff 范围内。
            # 同时, if i >= indexDiff: 语句保证了相同桶的索引差值始终在 indexDiff 范围内
            if bucket_num in all_buckets:
                return True

            # 3. 比较上一个桶
            if (bucket_num - 1) in all_buckets and abs(
                all_buckets[bucket_num - 1] - v
            ) <= valueDiff:
                return True

            # 4. 比较下一个桶
            if (bucket_num + 1) in all_buckets and abs(
                all_buckets[bucket_num + 1] - v
            ) <= valueDiff:
                return True

            # 5. 将值放进第bucket_num桶内
            all_buckets[bucket_num] = v  # 把v放入桶中

            # 6. 滑动窗口.
            # 如果不构成返回条件，那么当i >= indexDiff 的时候就要删除旧桶了，
            # 以维持桶中的元素索引跟下一个i+1索引只差不超过indexDiff
            # !!!: 保证了 all_buckets 的数量始终维持在 indexDiff 内。
            if i >= indexDiff:
                all_buckets.pop(nums[i - indexDiff] // bucket_size)

        return False


s = Solution()
res = s.containsNearbyAlmostDuplicate([1, 2, 3, 1], 3, 0)
print(res)
