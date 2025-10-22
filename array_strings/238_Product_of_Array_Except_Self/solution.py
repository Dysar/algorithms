class Solution(object):
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        res = [1] * len(nums)
        
        prefix=1 
        # O(n) time complexity
        for i in range(len(nums)): # iterate through the array, prefix is the product of the previous elements
            res[i] = prefix
            prefix *= nums[i]
        postfix = 1
        # O(n) time complexity
        for i in range(len(nums)-1, -1, -1): # start from the last index, go backwards
            res[i] *= postfix
            postfix *= nums[i]
        
        # O(1) mem complexity
        return res