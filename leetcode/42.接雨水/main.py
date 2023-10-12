# coding=utf-8


class Solution:
    def trap(self, height: List[int]) -> int:
        st = [0]
        n = len(height)
        if n <= 1:
            return 0

        area = 0
        for i in range(1, n):
            if height[i] < height[st[-1]]:
                st.append(i)
            elif height[i] == height[st[-1]]:
                st.append(i)
            else:
                while st and height[i] > height[st[-1]]:
                    mid = st.pop()
                    if st:
                        h = min(height[i], height[st[-1]]) - height[mid]
                        w = i - st[-1] - 1

                        area += h * w
                st.append(i)
            # print(area)

        return area


if __name__ == "__main__":
    # height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
    height = [4, 2, 0, 3, 2, 5]
    result = test(height)
    print(result)
