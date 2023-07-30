def is_palindrome_number_1(x: int) -> bool:
    return str(x) == str(x)[::-1]


def is_palindrome_number_2(x: int) -> bool:
    s = str(x)
    l = len(s)  # noqa: E741
    half = l // 2

    return s[:half] == s[-1 : -half - 1 : -1]


def is_palindrome_number_3(x: int) -> bool:
    """
    时间复杂度: O(log10(n))
    空间复杂度: O(1)
    """
    if x < 0 or (x % 10 == 0 and x != 0):
        return False

    reversed_number = 0
    while x > reversed_number:
        reversed_number = reversed_number * 10 + x % 10
        x //= 10

    print(reversed_number, x)
    return x == reversed_number or x == reversed_number // 10


if __name__ == "__main__":
    num1 = 1221
    num2 = 12345
    num3 = 1234321

    print("way1".center(40, "-"))
    print(is_palindrome_number_1(num1))
    print(is_palindrome_number_1(num2))
    print(is_palindrome_number_1(num3))

    print("way2".center(40, "-"))
    print(is_palindrome_number_2(num1))
    print(is_palindrome_number_2(num2))
    print(is_palindrome_number_2(num3))

    print("way3".center(40, "-"))
    print(is_palindrome_number_3(num1))
    print(is_palindrome_number_3(num2))
    print(is_palindrome_number_3(num3))
