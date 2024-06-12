package main

import "testing"

func Test_hasCycle(t *testing.T) {
	var (
		node1 = &ListNode{Val: 3}
		node2 = &ListNode{Val: 2}
		node3 = &ListNode{Val: 0}
		node4 = &ListNode{Val: -4}
	)
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	if !hasCycle(node1) {
		t.Error("expected true, got false")
	}
}
