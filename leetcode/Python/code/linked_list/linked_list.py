class ListNode:
    def __init__(self, x) -> None:
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head: "ListNode") -> "ListNode":
        pre, cur = None, head

        while cur:
            cur.next, pre, cur = pre, cur, cur.next

        return pre
