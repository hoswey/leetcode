import "sort"

func fourSum(nums []int, target int) [][]int {

    sort.Ints(nums)
    var ans [][]int

    for i := 0; i < len(nums) - 3; i++ {

        if i != 0 && nums[i] == nums[i-1] {
            continue
        }

        ts := threeSum(nums[i], nums[i+1:], target-nums[i])
        ans = append(ans, ts...)
    }
    return ans
}

func threeSum(fir int, nums []int, target int) [][]int {

    var ret [][]int
    for i := 0; i < len(nums) - 2; i++ {

        if i != 0 && nums[i] == nums[i-1] {
            continue
        }

        j := i + 1
        k := len(nums) - 1

        for j < k {         
            sum := nums[i] + nums[j] + nums[k]
            if sum < target {
                j++
                for j < k && nums[j] == nums[j-1] {
                    j++
                }
            } else if sum > target {
                k--
                for j < k && nums[k] == nums[k+1] {
                    k--
                }
            } else {
                ret = append(ret, []int{fir, nums[i], nums[j], nums[k]})
                j++
                for j < k && nums[j] == nums[j-1] {
                    j++
                }                
                k--
                for j < k && nums[k] == nums[k+1] {
                    k--
                }                
            }
        }
    }

    return ret
} 

