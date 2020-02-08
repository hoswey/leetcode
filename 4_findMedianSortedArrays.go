func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    l := len(nums1) + len(nums2)
    m := findKth(nums1, nums2, (l+1)/2)
    n := findKth(nums1, nums2, (l+2)/2)

    return float64(m+n) / 2
}


func findKth(nums1, nums2[] int, k int) int {

	if len(nums1) > len(nums2) {
		return findKth(nums2, nums1, k)
	}

	if len(nums1) == 0{
		return nums2[k-1]
	}

	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	m := min(len(nums1), k/2)

	if nums1[m-1] < nums2[m-1] {
		return findKth(nums1[m:], nums2, k - m)
	} else {
		return findKth(nums2[m:], nums1, k - m)
	}
}


func min(a, b int) int {

	if a < b {
		return a 
	} else {
		return b
	}
} 