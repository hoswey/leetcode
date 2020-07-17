//给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

//示例 1 :

//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//示例 2 :

//输入: 9973
//输出: 9973
//解释: 不需要交换。
package main

import (
  "fmt"
  "strconv"
)

func maximumSwap(num int) int {

  sNum := []byte(strconv.Itoa(num))

  arr := make([]int, 10)
  for i := range sNum {
    arr[sNum[i] - '0'] = i
  }

  for i, _ := range sNum {
    for j := 9; j >= 0; j-- {

      if arr[j] == 0 || arr[j] <= i || j <= int(sNum[i] - '0') {
        continue
      }

      sNum[i], sNum[arr[j]] = sNum[arr[j]], sNum[i]
      goto END
    }
  }

  END:

  ans, _ := strconv.Atoi(string(sNum))
  return ans
}

func main() {

  var input int

  //input = 2736
  //fmt.Printf("input:%d, ouput:%d, expected:%d\n", input, maximumSwap(input), 7236)

  //input = 9973
  //fmt.Printf("input:%d, ouput:%d, expected:%d\n", input, maximumSwap(input), 9973)

  //input = 10
  //fmt.Printf("input:%d, ouput:%d, expected:%d\n", input, maximumSwap(input), 10)

  input = 98368
  fmt.Printf("input:%d, ouput:%d, expected:%d\n", input, maximumSwap(input), 98863)


}
