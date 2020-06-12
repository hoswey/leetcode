//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

//请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

//你可以假设 nums1 和 nums2 不会同时为空。

//示例 1:

//nums1 = [1, 3]
//nums2 = [2]

//则中位数是 2.0
//示例 2:

//nums1 = [1, 2]
//nums2 = [3, 4]

//则中位数是 (2 + 3)/2 = 2.5
package main
import (
  "fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

  l := len(nums1) + len(nums2)

  a := findKth(nums1, nums2, (l+1)/2)
  b := findKth(nums1, nums2, (l+2)/2)

  return (float64(a) + float64(b))/2
}

func findKth(nums1 []int, nums2 []int, k int) int {

  fmt.Printf("nums1 %v, nums2: %v, k: %d\n", nums1, nums2, k)

  if len(nums1) > len(nums2) {
    return findKth(nums2, nums1, k)
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

  m := min(len(nums1), k/2)

  if nums1[m-1] < nums2[m-1] {
    return findKth(nums1[m:], nums2, k - m)
  } else {
    return findKth(nums1, nums2[m:], k - m)
  }

  return 0
}

func min(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}


func main() {

  var nums1, nums2 []int

  nums1 = []int{1,3}
  nums2 = []int{2}

  fmt.Printf("nums1: %v, nums2: %v, ouput: %v, expected: %v\n", nums1, nums2, findMedianSortedArrays(nums1, nums2), 2.0)

  nums1 = []int{1,2}
  nums2 = []int{3,4}

  fmt.Printf("nums1: %v, nums2: %v, ouput: %v, expected: %v\n", nums1, nums2, findMedianSortedArrays(nums1, nums2), 2.5)

  nums1 = []int{0,0}
  nums2 = []int{0,0}

  fmt.Printf("nums1: %v, nums2: %v, ouput: %v, expected: %v\n", nums1, nums2, findMedianSortedArrays(nums1, nums2), 0)

}
