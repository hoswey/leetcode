func permute(nums []int) [][]int {

    var ans [][]int
    helper(nil, nums, &ans)
    return ans
}

func helper(prefix []int, cur []int, ans *[][]int) {

    if len(cur) == 1 {
        arr := append(prefix, cur...)
        *ans = append(*ans, arr)
        return
    }

    for i := 0; i < len(cur); i++ {
        cur[0], cur[i] = cur[i], cur[0]
        prefixCopy := make([]int, len(prefix))
        copy(prefixCopy, prefix)
        prefixCopy = append(prefixCopy, cur[0])
        helper(prefixCopy, cur[1:], ans)
        cur[0], cur[i] = cur[i], cur[0]
    }

}
