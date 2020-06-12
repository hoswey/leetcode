// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

// 注意:
// 不能使用代码库中的排序函数来解决这道题。

// 示例:

// 输入: [2,0,2,1,1,0]
// 输出: [0,0,1,1,2,2]

func sortColors(nums []int)  {

	if len(nums) <= 1 {
		return
	}

	//find 1 as pivot
	var idxOfOne int
	for i := range nums {
		if nums[i] == 1 {
			idxOfOne = i
		}
	}

	nums[idxOfOne], nums[len(nums) - 1] = nums[len(nums) - 1], nums[idxOfOne]

	var j int 
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[len(nums) - 1] {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	nums[j], nums[len(nums)-1] = nums[len(nums)-1], nums[j]
}