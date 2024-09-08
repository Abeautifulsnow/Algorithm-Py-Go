from itertools import zip_longest


class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        word1_l, word2_l = len(word1), len(word2)
        i, j = 0, 0

        new_s = ""

        while i < word1_l or j < word2_l:
            if i < word1_l:
                new_s += word1[i]
                i += 1

            if j < word2_l:
                new_s += word2[j]
                j += 1

        return new_s

    def mergeAlternately1(self, word1: str, word2: str) -> str:
        new_s = ""

        for x, y in zip_longest(word1, word2):
            if x:
                new_s += x
            if y:
                new_s += y

        return new_s
