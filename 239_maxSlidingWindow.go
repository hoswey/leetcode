import "container/list"
//https://zhuanlan.zhihu.com/p/26432350
func maxSlidingWindow(nums []int, k int) []int {

    l := list.New()
    var ans []int

    for i := 0; i < len(nums); i++ {

        for l.Len() != 0 && nums[i] > nums[l.Back().Value.(int)] {
            l.Remove(l.Back())
        }
        l.PushBack(i)

        for l.Len() != 0 && l.Front().Value.(int) <= i - k {
            l.Remove(l.Front())
        }

        if i >= k - 1 {
            ans = append(ans, nums[l.Front().Value.(int)])
        }
    }

    return ans

}
