// 给你一个整数数组 arr 和两个整数 k 和 threshold 。
// 请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。

// 示例 1：
// 输入：arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
// 输出：3
// 解释：子数组 [2,5,5],[5,5,5] 和 [5,5,8] 的平均值分别为 4，5 和 6 。其他长度为 3 的子数组的平均值都小于 4 （threshold 的值)。

// 示例 2：
// 输入：arr = [1,1,1,1,1], k = 1, threshold = 0
// 输出：5

// 示例 3：
// 输入：arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
// 输出：6
// 解释：前 6 个长度为 3 的子数组平均值都大于 5 。注意平均值不是整数。

// 示例 4：
// 输入：arr = [7,7,7,7,7,7,7], k = 7, threshold = 7
// 输出：1
package main
import "fmt"

func numOfSubarrays(arr []int, k int, threshold int) int {

	if len(arr) < k {
		return 0
	}

	threshold = k * threshold

	var cur int
	for i := 0; i < k; i++ {
		cur += arr[i]
	}

	var ans int
	if cur >= threshold {
		ans++
	}

	for i := 1; i <= len(arr) - k; i++ {
		cur = cur - arr[i-1] + arr[i+k-1]
		if cur >= threshold{
			ans++
		}
	}
	return ans
}

func main() {

	var arr []int
	var k, threshold int

	arr = []int{2,2,2,2,5,5,5,8}
	k = 3
	threshold = 4
	fmt.Printf("arr:%v, k:%d, threshold:%d, output:%d, expected:%d\n", arr, k, threshold, numOfSubarrays(arr,k,threshold), 3)

	arr = []int{1,1,1,1,1}
	k = 1
	threshold = 0
	fmt.Printf("arr:%v, k:%d, threshold:%d, output:%d, expected:%d\n", arr, k, threshold, numOfSubarrays(arr,k,threshold), 5)

	arr = []int{11,13,17,23,29,31,7,5,2,3}
	k = 3
	threshold = 5
	fmt.Printf("arr:%v, k:%d, threshold:%d, output:%d, expected:%d\n", arr, k, threshold, numOfSubarrays(arr,k,threshold), 6)

	arr = []int{7,7,7,7,7,7,7}
	k = 7
	threshold = 7
	fmt.Printf("arr:%v, k:%d, threshold:%d, output:%d, expected:%d\n", arr, k, threshold, numOfSubarrays(arr,k,threshold), 1)	
}

