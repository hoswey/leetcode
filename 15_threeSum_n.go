package main

import (
    "fmt"
    "sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

//  

// 示例：

// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

func threeSum(nums []int) [][]int {

    sort.Ints(nums)
    
    var ans [][]int
    for i := 0; i < len(nums) - 2; i++ {

        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        j := i + 1
        k := len(nums) - 1
        for j < k {

            sum := nums[i] + nums[j] + nums[k]

            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})

                j++
                for j < k && nums[j] == nums[j-1] {
                    j++
                }

                k--
                for k > j && nums[k] == nums[k+1] {
                    k--
                }

            } else if sum < 0 {

                j++
                for j < k && nums[j] == nums[j-1] {
                    j++
                }                

            } else {
                k--
                for k > j && nums[k] == nums[k+1] {
                    k--
                }
            }

        }
    }

    return ans
}

func main() {

    input := []int{-1, 0, 1, 2, -1, -4}
    fmt.Printf("%v\n", threeSum(input))

    input = []int{-1, 0}
    fmt.Printf("%v\n", threeSum(input))

    fmt.Printf("%v\n", threeSum(nil))
}