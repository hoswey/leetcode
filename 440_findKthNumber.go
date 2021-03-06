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

	//前缀
	prefix := 1

	//作为一个指针，指向当前所在位置，当p==k时，也就是到了排位第k的数	
	p := 1
	for p < k {

		//获得当前前缀下所有子节点的和
		count := getCount(prefix, n)

		fmt.Printf("p:%d, prefix:%d, count:%d\n", p, prefix, count)
		//第k个数在当前前缀下		
		if p + count > k {
			p++
			prefix *= 10
		} else {
			p += count
			prefix++
		}

	}

	return prefix
    
}

func getCount(i, n int) int {

	cur := i
	next := i+1

	var count int
	for cur <= n {
		count += min(next, n+1) - cur
		cur *= 10
		next *= 10
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	
	var n, k int
	n = 13
	k = 2
	fmt.Printf("n:%d, k:%d, ouput:%d, expected:%d\n", n, k, findKthNumber(n, k), 10)

	k = 5
	fmt.Printf("n:%d, k:%d, ouput:%d, expected:%d\n", n, k, findKthNumber(n, k), 13)

}