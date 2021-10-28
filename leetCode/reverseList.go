package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {

	node2 := ListNode{Val: 3, Next: nil}
	node1 := ListNode{Val: 2, Next: &node2}
	node := ListNode{1, &node1}
	println(node.Next.Val)
	list := reverseList(&node)
	println(list.Val, list.Next.Val)
}
