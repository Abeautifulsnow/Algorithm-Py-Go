class Solution:
    def myAtoi(self, s: str) -> int:
        start, end = None, None
        signed = "-+"
        str_digit = "0123456789"
        s_length = len(s)

        # 定义数值的最大和最小边界
        min_i, max_i = pow(-2, 31), pow(2, 31)

        # 定义字符串的长度
        if s_length > 200:
            return 0

        for i in range(s_length):
            # 1. 如果刚开始的字符串是空格，就继续循环
            if s[i] == " ":
                continue

            # 2. 首字符是+-或者是数字，且start没有被赋值，就进入这个逻辑
            # 获取字符串切片的初值
            elif (s[i] in signed or s[i] in str_digit) and start is None:
                start = i
                end = i

            # 3. 当start被赋值后，进入下面逻辑
            elif start is not None and start >= 0:
                # 判断是否是字符串，若是，给定end的值
                if i - start == 1 and s[i] in str_digit:
                    end = i
                # 当end有值后，判断后一位字符是不是数字，若是，指针后移
                if end and i - end == 1 and s[i] in str_digit:
                    end = i
                # 若end存在，且后一个字符不是数字，就跳出循环
                if end and s[i] not in str_digit:
                    break
            else:
                break

        if end is not None and end >= 0:
            res = s[start : end + 1]

            # 如果只拿到一个+, 返回0
            if res in signed:
                return 0

            # 下边界
            if int(res) < min_i:
                return min_i

            # 上边界
            if int(res) >= max_i:
                return max_i - 1

            return int(res)
        else:
            return 0


s = Solution()
print(s.myAtoi(".2"))
