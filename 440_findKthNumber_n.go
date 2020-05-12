// 给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。

// 注意：1 ≤ k ≤ n ≤ 109。

// 示例 :

// 输入:
// n: 13   k: 2

// 输出:
// 10

// 解释:
// 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
 



package main
import "fmt"

func findKthNumber(n int, k int) int {

	p := 1
	prefix := 1

	for p < k {

		count := getCount(prefix, n)

		if p + count > k {

			p += 1
			prefix *= 10

		} else {
			p += count
			prefix += 1
		}
	}

    return prefix
}

func getCount(prefix, n int) int{

	var count int
	next := prefix + 1
	for prefix < n {
		count += min(n+1, next) - prefix

		prefix *= 10
		next *= 10
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	
	var n, k int
	n = 13
	k = 2
	fmt.Printf("n:%d, k:%d, ouput:%d, expected:%d\n", n, k, findKthNumber(n, k), 10)

	k = 5
	fmt.Printf("n:%d, k:%d, ouput:%d, expected:%d\n", n, k, findKthNumber(n, k), 13)

}