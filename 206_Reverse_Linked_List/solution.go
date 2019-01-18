/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    node := head
    var prev *ListNode
    
    for node != nil {
        next := node.Next
        node.Next = prev
        prev = node
        node = next
    }
    
    return prev
}
