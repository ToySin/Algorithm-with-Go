package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, m+n)

	for idx1, idx2 := 0, 0; idx1 < m || idx2 < n; {
		switch {
		case idx1 == m:
			result[idx1+idx2] = nums2[idx2]
			idx2++
		case idx2 == n:
			result[idx1+idx2] = nums1[idx1]
			idx1++
		case nums1[idx1] <= nums2[idx2]:
			result[idx1+idx2] = nums1[idx1]
			idx1++
		case nums1[idx1] > nums2[idx2]:
			result[idx1+idx2] = nums2[idx2]
			idx2++
		}
	}
	for idx, _ := range nums1 {
		nums1[idx] = result[idx]
	}
}
