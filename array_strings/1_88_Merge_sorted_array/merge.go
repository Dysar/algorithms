package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	for k := m + n - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i] // we set the kth element to num 1 if num 1 is greater than num 2
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j] // we set the kth element to num 2 if num 2 is greater than num 1
			j--
		}
	}
}

// exec plan
// k is 5 (index of last element of nums1
// nums [1,2,3], nums2 [2,5,6
// if nums1 el > nums2 el => if 3 > 6 false
// nums2 el > nums1 el => [      6]. i=2; j=1

// next iter: if 3 > 5 false; [     5,6]. i=2 j=0 k=4
// next iter: if 3 > 2 true; [     3,5,6]. i=1, j=0 k=3
// next iter: if 2 > 2 false; [   2,3,5,6]. i=1 j=-1 k=2
// next iter: j <= 0 so we skip and keep elements at k=1	[ 2,2,3,5,6]
// next iter: j <= 0 so we skip and keep elements at k=0	[1,2,2,3,5,6]
