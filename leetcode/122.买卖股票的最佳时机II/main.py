# coding=utf-8

from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0

        for i in range(len(prices) - 1):
            tmp_p = prices[i + 1] - prices[i]

            if tmp_p > 0:
                profit += tmp_p

        return profit


# 创建堆（调整堆）
def heapify(arr, n, i):
    largest = i
    l = 2 * i + 1  # left = 2*i + 1
    r = 2 * i + 2  # right = 2*i + 2

    if l < n and arr[i] < arr[l]:
        largest = l

    if r < n and arr[largest] < arr[r]:
        largest = r

    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)


# 堆排序
def heapSort(arr):
    n = len(arr)

    # Build a maxheap.
    for i in range(n, -1, -1):
        heapify(arr, n, i)

    # 一个个交换元素
    for i in range(n - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify(arr, i, 0)
    return arr


if __name__ == "__main__":
    # prices = [7, 1, 5, 3, 6, 4]
    prices = [10, 3, 1, 2, 0, 8, 4, 3, 9]
    # print(Solution().maxProfit(prices))
    print(heapSort(prices))
