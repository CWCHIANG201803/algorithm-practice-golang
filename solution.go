package solution

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode = nil
	var cur *ListNode = head
	var nxt *ListNode = nil

	for cur != nil {
		nxt = cur.Next
		cur.Next = prev
		prev = cur

		cur = nxt
	}
	head = prev
	return head
}
