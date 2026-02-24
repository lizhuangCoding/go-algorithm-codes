package main

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	demo := head
	for demo != nil {
		demo = demo.Next
		n++
	}

}
