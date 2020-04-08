package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Using two loops
// Time complexity: O(n)
// Space complexity: O(n)
func middleNodeLoop(head *ListNode) *ListNode {
	lght := 0
	midIdx := 0

	for cur := head; cur != nil; cur = cur.Next {
		lght++
	}

	midIdx = lght / 2

	i := 0
	for cur := head; cur != nil; cur = cur.Next {
		if i == midIdx {
			return cur
		}

		i++
	}

	return nil
}

// Output to a slice
// Time complexity: O(n)
// Space complexity: O(n)
func middleNodeArray(head *ListNode) *ListNode {
	a := make([]*ListNode, 0, 100)

	for cur := head; cur != nil; cur = cur.Next {
		a = append(a, cur)
	}

	return a[len(a)/2]
}

// Fast and slow pointer
// Time complexity: O(n)
// Space complexity: O(1)
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
