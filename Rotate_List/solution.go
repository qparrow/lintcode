package Rotate_List


type ListNode struct {
	Val int
	Next *ListNode
}


func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0  {
		return head
	}
	end := head
	length := 1
	for end.Next != nil {
		length ++
		end = end.Next
	}

	if (k % length) == 0 {
		return head
	}

	end.Next = head
	for i := 0; i< (length-(k % length)); i++ {
		end = end.Next
	}
	head = end.Next
	end.Next = nil
	return head
}