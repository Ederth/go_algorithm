package week1

func main21() {
	a1 := []int{1, 2, 4}
	a2 := []int{1, 3, 4}

	l1 := arr2List(a1)
	l2 := arr2List(a2)

	mergeTwoLists(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	h := l
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			l.Next = &ListNode{Val: l1.Val}
			l = l.Next
			l1 = l1.Next
		} else {
			l.Next = &ListNode{Val: l2.Val}
			l = l.Next
			l2 = l2.Next
		}
	}

	return h.Next
}

func arr2List(arr []int) *ListNode {
	var head, last *ListNode
	for i, v := range arr {
		l := &ListNode{
			Val:  v,
			Next: nil,
		}

		if i == 0 {
			head = l
			last = l
			continue
		}

		last.Next = l
		last = l
	}

	return head
}
