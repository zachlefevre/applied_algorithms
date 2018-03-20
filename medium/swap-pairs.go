// Solution to https://leetcode.com/problems/swap-nodes-in-pairs
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (lst *ListNode) append(apnd *ListNode) {
	for lst.Next != nil {
		lst = lst.Next
	}
	lst.Next = apnd
}
func (lst *ListNode) printList() {
	for n := lst; n != nil; n = n.Next {
		fmt.Print(n.Val, "->")
	}
	fmt.Println("//")
}
func swap(prev, cur, nxt *ListNode) *ListNode {
	if prev != nil {
		prev.Next = nxt
	}
	cur.Next = nxt.Next
	nxt.Next = cur
	return nxt
}
func swapPairs(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil && cur.Next != nil {
		cur = swap(prev, cur, cur.Next)
		if prev == nil {
			head = cur
		}
		prev = cur.Next
		cur = prev.Next
	}
	return head
}

func main() {
	lst := &ListNode{Val: 0, Next: nil}
	for i := 1; i <= 11; i++ {
		lst.append(&ListNode{Val: i, Next: nil})
	}
	lst.printList()
	lst = swapPairs(lst)
	lst.printList()
}
