package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur *ListNode = head

	for cur != nil {
		var tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func reverseListRecursion(head *ListNode) *ListNode {
	// 1. 递归终止条件
	if head == nil || head.Next == nil {
		return head
	}
	var p = reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
