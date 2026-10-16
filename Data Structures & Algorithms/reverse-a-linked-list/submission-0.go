/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}
// func reverseList(head *ListNode) *ListNode {
//     var prev *ListNode 

// 	for head != nil {
// 		head.Next = prev
// 		prev = head
// 		head = head.Next
// 	}

// 	return prev
// }
