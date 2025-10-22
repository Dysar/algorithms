# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    """
        :type root: Optional[TreeNode]
        :rtype: List[int]
    """
    def inorderTraversal(self, root):
        res = []

        def inorder(root):
            if not root: # nill
                return
            inorder(root.left)
            res.append(root.val) #process root node itself
            inorder(root.right)
        
        inorder(root)
        return res
        
        # 1 -> traverse entire left subtree -> traverse entire right subtree. [1,3,2] O(n). mem O(n)



        