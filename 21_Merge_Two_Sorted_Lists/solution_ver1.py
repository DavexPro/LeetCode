# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        # Boundary check
        if l1 is None:
            return l2
        elif l2 is None:
            return l1
        
        # Recursive
        node_head = None
        
        if l1.val < l2.val:
            node_head = l1
            node_head.next = self.mergeTwoLists(l1.next, l2)
        else:
            node_head = l2
            node_head.next = self.mergeTwoLists(l2.next, l1)
        
        return node_head
