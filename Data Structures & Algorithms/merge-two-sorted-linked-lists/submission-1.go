import "slices"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	
	listElemts := make([]int, 0, 100)

	for list1 != nil {
		listElemts = append(listElemts, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		listElemts = append(listElemts, list2.Val)
		list2 = list2.Next
	}

	slices.Sort(listElemts)

	sortedList := &ListNode{}
	curr := sortedList 

	for _, val := range listElemts {
    	curr.Next = &ListNode{Val: val} 
    	curr = curr.Next          
	}

	return sortedList.Next
}
