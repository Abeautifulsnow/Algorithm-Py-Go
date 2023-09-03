# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth_with_dfs(self, root: Optional[TreeNode]) -> int:
        """深度检索"""
        # 如果没有值, 那就返回0
        if not root:
            return 0

        # 沿着左树 或者 右树, 一直找, 谁最大返回谁
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))

    def maxDepth_with_bfs(self, root: Optional[TreeNode]) -> int:
        """宽度检索"""
        if not root:
            return 0

        result = 0
        node_list = [root]
        while node_list:
            result += 1
            n = len(node_list)
            for _ in range(n):
                if node_list[0].left:
                    node_list.append(node_list[0].left)
                if node_list[0].right:
                    node_list.append(node_list[0].right)
                node_list.pop(0)

        return result
