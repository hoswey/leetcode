//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

//上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

//示例:

//输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
package main

import (
  "fmt"
)

func maxArea(height []int) int {

  if len(height) <= 2 {
    return 0
  }

  maxLeft := make([]int, len(height))
  maxRight := make([]int, len(height))

  for i := 1; i < len(maxLeft); i++ {
    maxLeft[i] = max(maxLeft[i-1], height[i-1])
  }

  for i := len(maxRight) - 2; i >= 0; i-- {
    maxRight[i] = max(maxRight[i+1], height[i+1])
  }

  var ans int
  for i := 1; i < len(height) - 1; i++ {

    lo := min(maxLeft[i], maxRight[i])

    if lo < height[i] {
      continue
    }
    ans += lo - height[i]
  }

  return ans
}

func max(a, b int) int {
  if a > b {
    return a
  }

  return b
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func main() {

  input := []int{0,1,0,2,1,0,1,3,2,1,2,1}
  fmt.Printf("input: %v, ouput: %v, expected: %d\n", input, maxArea(input), 6)
}
