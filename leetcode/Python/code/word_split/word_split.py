"""
题目大致如此, 给定一个函数 split, 入参分别为字符串 text 和 一个单词集合 dic(按照字符串的顺序排列的单词顺序), 
然后输出 text 根据 dic 里面的单词切割出来的数组, 比如:

例子1:

字符串text为: 'appleddesktop'
单词集合dict为: set(['a', 'app', 'apple', 'led', 'desk', 'top'])
输出: ['app', 'led', 'desk', 'top']
例子2:

字符串text为: 'appledesktop'
单词集合dict为: set(['a', 'app', 'apple', 'led', 'desk', 'top'])
输出: ['apple', 'desk', 'top']
「可以发现当输出的数组再转为字符串的时候, 和输入的 text 一模一样且字符顺序一致。」
请设计代码实现如上所述的逻辑, 不限语言。
"""

# 单词列表必须有顺序才行
def split_word_in_order(text: str, word_dic: set | list) -> list:
    if isinstance(word_dic, set):
        word_dic = [word for word in word_dic]

    length = len(word_dic)

    while length:
        result = []
        temp = text

        # 遍历单词列表
        for i in range(length):
            word = word_dic[i]
            if word[0] == temp[0] and word in temp:
                result.append(word)
                # 将字符串切片，然后继续遍历与单词列表比较
                temp = temp.split(word)[-1]

        # 理想情况下，字典列表和字符串完全匹配
        if "".join(result) == text:
            return result
        else:
            # 否则就将单词列表的第一个元素给剔除掉，重新继续查找
            word_dic.pop(0)
            length = len(word_dic)


# 通用版，单词列表可以无序
def split_word_disordered(text: str, word_dic: set | list) -> list | None:
    length = len(text)
    dp = [False] * length
    dict_c = {}

    for word in word_dic:
        dict_c[word] = True

    for i in range(length):
        str_var = ""

        for j in range(i, length):
            str_var += text[j]

            if dict_c.get(str_var) and (i == 0 or dp[i - 1]):
                if dp[i - 1]:
                    dp[j] = []
                    dp[j].extend([*dp[i - 1], str_var])
                else:
                    dp[j] = [str_var]

                if j == length - 1:
                    return dp[j]

    return None


if __name__ == "__main__":
    t = "appleddesktop"
    w = set(["a", "led", "app", "apple", "desk", "top"])
    # print(split_word_in_order(t, w))
    print(split_word_disordered(t, w))
