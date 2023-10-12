# coding=utf-8


def test(nums):
    n = len(nums)
    if n == 1:
        return nums[0]

    dp = [0 for _ in range(n)]
    ##1.dp[i] 偷窃i号房屋获得最大价值
    ##2. 递推公式
    ##3. 初始化
    ##遍历顺序
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    for i in range(2, n):
        dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])

    return dp[n - 1]


if __name__ == "__main__":
    # nums = [1, 2, 3, 1]
    nums = [2, 7, 9, 3, 1]
    result = test(nums)
    print(result)
