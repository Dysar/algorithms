import unittest
import os
import sys

# Ensure this directory is on sys.path so `solution` can be imported when running from project root
sys.path.insert(0, os.path.dirname(__file__))

from solution import Solution


class TestProductExceptSelf(unittest.TestCase):
    def setUp(self):
        self.s = Solution()

    def test_basic(self):
        self.assertEqual(self.s.productExceptSelf([1, 2, 3, 4]), [24, 12, 8, 6])

    def test_with_zero_single(self):
        # One zero -> only the index of zero gets product of others
        self.assertEqual(self.s.productExceptSelf([0, 1, 2, 3, 4]), [24, 0, 0, 0, 0])

    def test_with_zeros_two(self):
        # Two zeros -> all zeros
        self.assertEqual(self.s.productExceptSelf([0, 1, 0, 3, 4]), [0, 0, 0, 0, 0])

    def test_negatives(self):
        self.assertEqual(self.s.productExceptSelf([-1, 1, 0, -3, 3]), [0, 0, 9, 0, 0])

    def test_single_element(self):
        # Depending on specification, usually single element case returns [1]
        self.assertEqual(self.s.productExceptSelf([5]), [1])


if __name__ == '__main__':
    unittest.main()
