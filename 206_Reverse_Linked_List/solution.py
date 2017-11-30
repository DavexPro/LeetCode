# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        # Boundary Checl
        if not head:
            return None
        
        # Iteratively
        node_head = None
        node_prev = None
        node_current = head
        
        while node_current:
            
            node_next = node_current.next
            
            if not node_next:
                node_head = node_current
                
            node_current.next = node_prev
            
            node_prev = node_current
            node_current = node_next
            
        return node_head
