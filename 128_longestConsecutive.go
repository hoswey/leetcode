//给定一个未排序的整数数组，找出最长连续序列的长度。

//要求算法的时间复杂度为 O(n)。

//示例:

//输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4
package main

import "fmt"

func longestConsecutive(nums []int) int {

  if len(nums) == 0 {
    return 0
  }

    m := make(map[int]int)

    var ans int

    for _, num := range nums {
      m[num] = num
    }

    for _, num := range nums {

      i := 1

      for {
        if _, ok := m[num+i]; ok {
          i++
        } else {
          break
        }
      }

      if i > ans {
        ans = i
      }
    }

    return ans
}

func main() {
  var input []int

  input = []int{100,4,200,1,3,2}

  fmt.Printf("input:%v, output: %v, expected: %v\n", input, longestConsecutive(input), 4)
}

