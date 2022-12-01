"""binary_search.py 二分查找
"""


def binary_search(order_arr: list[int], search: int):
    if len(order_arr) == 0 or order_arr is None:
        return -1
    start, end = 0, len(order_arr) - 1
    search_idx = 0

    while start <= end:
        mid = (end + start) // 2

        if search == order_arr[mid]:
            search_idx = mid
            break

        if search < order_arr[mid]:
            end = mid - 1

        if search > order_arr[mid]:
            start = mid + 1

    return search_idx


# demo: [2,5,6,8,12,15,17,23,27,31,39,40,45,56,79,90]
arr = [2, 5, 6, 8, 12, 15, 17, 23, 27, 31, 39, 40, 45, 56, 79, 90]
print(binary_search(arr, 39))
