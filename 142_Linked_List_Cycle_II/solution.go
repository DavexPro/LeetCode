/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    linkLen := lenCycle(head)
    if linkLen == -1 {
        return nil
    }
    
    slow := head
    fast := head
    
    for i := 0; i < linkLen; i++ {
        fast = fast.Next
    }
    
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }
    
    return slow
}

func lenCycle(head *ListNode) int {
    if head == nil {
        return -1
    }
    
    fast := head.Next
    slow := head
    
    for fast != slow {
        if fast == nil || fast.Next == nil {
            return -1
        }
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    linkLen := 1
    slow = slow.Next
    
    for slow != fast {
        slow = slow.Next
        linkLen ++
    }
    
    return linkLen
}
