/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	s := len(nums1) + len(nums2)

	n1 := findTheKth(nums1, nums2, (s+1)/2)
	n2 := findTheKth(nums1, nums2, (s+2)/2)
	return (float64(n1) + float64(n2))/2
}

func findTheKth(nums1 []int, nums2 []int, k int) int {

	if len(nums1) > len(nums2) {
		return findTheKth(nums2, nums1, k)
	}

	if len(nums1) == 0 {
		return nums2[k-1]
	}

	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	m := k/2
	if len(nums1) < m {
		m = len(nums1)
	}
	if nums1[m-1] < nums2[m-1] {
		return findTheKth(nums1[m:], nums2, k - m)
	} else {
		return findTheKth(nums1, nums2[m:], k - m)
	}
}
// @lc code=end

