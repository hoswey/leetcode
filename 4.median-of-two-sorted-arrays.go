/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l := len(nums1) + len(nums2)
	n1 := findKth(nums1, nums2, (l + 1)/2)
	n2 := findKth(nums1, nums2, (l + 2)/2)

	return float64(n1+n2)/2
}

func findKth(nums1, nums2 []int, k int) int {

	if len(nums1) > len(nums2) {
		return findKth(nums2, nums1, k)
	}

	if len(nums1) == 0 {
		return nums2[k - 1]
	}

	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	h := k / 2
	if h > len(nums1) {
		h = len(nums1)
	}

	if nums1[h-1] < nums2[h-1] {
		return findKth(nums1[h:], nums2, k - h)
	} else {
		return findKth(nums1, nums2[h:], k - h)
	}
}
// @lc code=end

