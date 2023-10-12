from collections import deque
from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []

        phone = {
            "2": ["a", "b", "c"],
            "3": ["d", "e", "f"],
            "4": ["g", "h", "i"],
            "5": ["j", "k", "l"],
            "6": ["m", "n", "o"],
            "7": ["p", "q", "r", "s"],
            "8": ["t", "u", "v"],
            "9": ["w", "x", "y", "z"],
        }

        def back_track(combination: str, rest_digit: List[str]):
            # 执行结果收集
            if rest_digit == "":
                result.append(combination)
            else:
                # 执行字符串的拼接
                for letter in phone[rest_digit[0]]:
                    back_track(combination + letter, rest_digit[1:])

        result = []
        back_track("", digits)

        return result

    def letterCombinations1(self, digits: str) -> List[str]:
        if not digits:
            return []
        phone = ["abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]
        queue = deque([""])  # 初始化队列
        for digit in digits:
            for _ in range(len(queue)):
                tmp = queue.popleft()
                for letter in phone[ord(digit) - 50]:  # 这里我们不使用 int() 转换字符串，使用ASCII码
                    queue.append(tmp + letter)

        # Convert deque to list.
        return list(queue)


if __name__ == "__main__":
    print(Solution().letterCombinations1("24"))
