import unittest
import os
import sys

# Ensure this directory is on sys.path so `solution` can be imported when running from project root
sys.path.insert(0, os.path.dirname(__file__))

from solution import Solution


class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class TestInorderTraversal(unittest.TestCase):
    def setUp(self):
        self.s = Solution()

    def test_empty(self):
        self.assertEqual(self.s.inorderTraversal(None), [])

    def test_single_node(self):
        root = TreeNode(1)
        self.assertEqual(self.s.inorderTraversal(root), [1])

    def test_left_skewed(self):
        # 3 -> 2 -> 1
        root = TreeNode(3, TreeNode(2, TreeNode(1)))
        self.assertEqual(self.s.inorderTraversal(root), [1, 2, 3])

    def test_right_skewed(self):
        # 1 -> 2 -> 3
        root = TreeNode(1, None, TreeNode(2, None, TreeNode(3)))
        self.assertEqual(self.s.inorderTraversal(root), [1, 2, 3])

    def test_mixed(self):
        #   1
        #    \
        #     2
        #    /
        #   3
        # Expected inorder: [1, 3, 2]
        root = TreeNode(1, None, TreeNode(2, TreeNode(3), None))
        self.assertEqual(self.s.inorderTraversal(root), [1, 3, 2])

    def test_full_balanced(self):
        #        4
        #      /   \
        #     2     6
        #    / \   / \
        #   1   3 5   7
        root = TreeNode(
            4,
            TreeNode(2, TreeNode(1), TreeNode(3)),
            TreeNode(6, TreeNode(5), TreeNode(7)),
        )
        self.assertEqual(self.s.inorderTraversal(root), [1, 2, 3, 4, 5, 6, 7])


if __name__ == '__main__':
    unittest.main()
