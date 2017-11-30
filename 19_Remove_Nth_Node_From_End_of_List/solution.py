# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        # Boundary Check
        if head == None or n == 0:
            return None
 
        # Iteratively
        node_ahead = head
        node_behind = head

        for _ in range(n):
            # Check N whether it's valid
            if not node_ahead:
                return None
            node_ahead = node_ahead.next
        
        if not node_ahead:
            return head.next
        
        while node_ahead.next:
            node_ahead = node_ahead.next
            node_behind = node_behind.next
        
        node_behind.next = node_behind.next.next
        
        # 返回倒数第 K 个节点
        return node_behind.next

        # 删除节点
        return head
