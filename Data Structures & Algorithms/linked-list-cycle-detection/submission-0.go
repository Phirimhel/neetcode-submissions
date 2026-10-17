func hasCycle(head *ListNode) bool {
	
	seen := make(map[*ListNode]struct{})
	
	for current := head; current != nil; current = current.Next {
		if _, ok := seen[current]; ok {
			return true
		} else {
			seen[current] = struct{}{}
		}
	} 

	return false 
}
