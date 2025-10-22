# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        if head is None:
            return None

        current = head
        # if current != nil && current.Next != nil -> 
        while current is not None and current.next is not None:
            # skip duplicate node
            if current.val == current.next.val:
                # skip duplicate node
                current.next = current.next.next
            else:
                current = current.next
        return head
        
        
        