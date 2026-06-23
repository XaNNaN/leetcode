package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, current, longer, min_node *ListNode
	// fmt.Printf("head: %v, current: %v, list1: %v, list2: %v\n", head, current, list1, list2)
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			min_node = list1
			list1 = list1.Next
		} else {
			min_node = list2
			list2 = list2.Next
		}
		if head == nil {
			head = min_node
			current = head
		} else {
			current.Next = min_node
			current = current.Next
		}
		// fmt.Printf("head: %v, current: %v, list1: %v, list2: %v\n", head, current, list1, list2)
	}
	if list1 != nil {
		longer = list1
	} else {
		longer = list2
	}
	if head == nil {
		head = longer
	} else {
		current.Next = longer
	}
	return head
}
