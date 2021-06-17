func search(nums []int, target int) int {

    if len(nums) == 0 {
        return -1
    }

    //二分查找
    //对于每个二分的值 nums[mid]，
    //假如 nums[lo] < nums[mid], 不含旋转
    //   且target < nums[mid], 向前
    //假如 nums[lo] > nums[mid] , 发生旋转
    //   假如target >= nums[lo], 向前
    //   假如target <  nums[mid], 向前

    lo := 0
    hi := len(nums) - 1

    for lo < hi {

        mid := lo + (hi - lo)/2

        fmt.Printf("1 lo=%d, hi=%d,mid=%d\n", lo, hi, mid)
        if nums[mid] == target {
            return mid
        }

        if nums[lo] <= target && target < nums[mid] || nums[mid] < nums[lo] && (target >= nums[lo] || target < nums[mid]) {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }

    if nums[lo] == target {
        return lo
    } else {
        return -1
    }

}
