# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        # Boundary Check
        if not head:
            return None
        
        # First Check Whether it is cirle LinkedList
        node_cycle = self.checkCycle(head)
        if not node_cycle:
            return None
        
        # Get the num of Cycle Nodes
        num_cycle = 1
        node_next = node_cycle.next
        while node_next != node_cycle:
            node_next = node_next.next
            num_cycle += 1
            
        # Use two pointer
        fast = head
        slow = head
        
        for _ in range(num_cycle):
            fast = fast.next
            
        while fast != slow:
             fast = fast.next
             slow = slow.next
        
        return slow
        
        
    def checkCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        # Boundary Check
        if not head:
            return None
        
        # Iteratively
        fast = head
        slow = head
        
        while fast:
            fast = fast.next
            
            if not fast:
                return None
            
            fast = fast.next
            slow = slow.next
            
            if fast == slow:
                return slow
        
        return None
            
        
