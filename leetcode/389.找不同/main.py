from collections import Counter
from functools import reduce
from operator import xor


class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        return chr(reduce(xor, map(ord, s + t)))

    def findTheDifference1(self, s: str, t: str) -> str:
        return list(Counter(t) - Counter(s))[0]

    def findTheDifference2(self, s: str, t: str) -> str:
        res = ord(t[-1])
        for i in range(len(s)):
            res += ord(t[i]) - ord(s[i])
        return chr(res)

    def findTheDifference3(self, s: str, t: str) -> str:
        diff = 0
        for c in t:
            diff += ord(c)
        for c in s:
            diff -= ord(c)
        return chr(diff)
