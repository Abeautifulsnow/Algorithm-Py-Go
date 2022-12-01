"""检查数据连续性
In: e.g.[1, 2, 3, 5, 7]
Out: e.g.['1-3', 5, 7]
"""


def check_arr_continurity(in_arr: list[int]):
    start = 0
    result = []

    for idx, el in enumerate(in_arr):
        if idx == 0:
            if el + 1 == in_arr[idx + 1]:
                start = el
            else:
                result.append(el)
        elif 0 < idx < len(in_arr) - 1:
            if el - 1 != in_arr[idx - 1] and el + 1 == in_arr[idx + 1]:
                start = el
            if el - 1 == in_arr[idx - 1] and el + 1 != in_arr[idx + 1]:
                result.append(f"{start}-{el}")
            if el - 1 != in_arr[idx - 1] and el + 1 != in_arr[idx + 1]:
                result.append(el)
        elif idx == len(in_arr) - 1:
            if el - 1 == in_arr[idx - 1]:
                result.append(f"{start}-{el}")
            else:
                result.append(el)

    return result


print(check_arr_continurity([2, 5, 7, 11, 19, 23, 24, 25, 27, 28, 29, 30]))
# Output: [2, 5, 7, 11, 19, '23-25', '27-30']
