class Solution:
    mapping = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    def romanToInt(self, s: str) -> int:
        s_len = len(s)
        sum = 0
        # jump表示IV，IX等情况的低位下标 => I所处的位置
        # 如果是I(低位)的下标，则跳过该次循环
        jump = -1

        for i in range(s_len - 1, -1, -1):
            # jump表示IV，IX等情况的低位下标 => I所处的位置
            # 如果是I(低位)的下标，则跳过该次循环
            if jump == i:
                jump = -1
                continue

            if i >= 1:
                if self.mapping[s[i]] > self.mapping[s[i - 1]]:
                    res = self.mapping[s[i]] - self.mapping[s[i - 1]]
                    sum += res
                    jump = i - 1
                else:
                    sum += self.mapping[s[i]]
            else:
                sum += self.mapping[s[i]]

        return sum


s = Solution()
print(s.romanToInt("MCMXCIV"))
