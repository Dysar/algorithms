import unittest
import os
import sys

# Ensure this directory is on sys.path so `solution` can be imported
sys.path.insert(0, os.path.dirname(__file__))

from solution import Solution, ListNode


def build_list(values):
    dummy = ListNode(0)
    current = dummy
    for v in values:
        current.next = ListNode(v)
        current = current.next
    return dummy.next


def to_list(head):
    result = []
    current = head
    while current is not None:
        result.append(current.val)
        current = current.next
    return result


class TestDeleteDuplicates(unittest.TestCase):
    def setUp(self):
        self.s = Solution()

    def test_empty(self):
        self.assertIsNone(self.s.deleteDuplicates(None))

    def test_single(self):
        head = build_list([1])
        self.assertEqual(to_list(self.s.deleteDuplicates(head)), [1])

    def test_no_duplicates(self):
        head = build_list([1, 2, 3])
        self.assertEqual(to_list(self.s.deleteDuplicates(head)), [1, 2, 3])

    def test_all_duplicates(self):
        head = build_list([2, 2, 2, 2])
        self.assertEqual(to_list(self.s.deleteDuplicates(head)), [2])

    def test_mixed_duplicates(self):
        head = build_list([1, 1, 2, 3, 3])
        self.assertEqual(to_list(self.s.deleteDuplicates(head)), [1, 2, 3])

    def test_tail_duplicates(self):
        head = build_list([1, 2, 3, 3, 3])
        self.assertEqual(to_list(self.s.deleteDuplicates(head)), [1, 2, 3])

    def test_head_duplicates(self):
        head = build_list([1, 1, 1, 2, 3])
        self.assertEqual(to_list(self.s.deleteDuplicates(head)), [1, 2, 3])


if __name__ == '__main__':
    unittest.main()
