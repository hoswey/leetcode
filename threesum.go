func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	var ret [][]int
	for k := 0; k < len(nums) - 2; k ++ {

		i := k + 1
		j := len(nums)
		
		var iUp, jDown bool
		var valueI, valueJ int
		for i < j {

			if iUp && nums[i] == valueI {
				i++
				iUp = true
				continue
			}

			if jDown && nums[j] == valueJ {
				j--
				jDown = true
				continue
			}

			iUp, jDown = true, true

			valueI = nums[i]
			valueJ = nums[j]
			sum := nums[k] + valueI + valueJ

			if sum == 0 {
				ret = append(ret, []int{nums[k], valueI, valueJ})
				i++
				j--
				iUp, jDown = true, true
			} else if (sum < 0) {
				i++
				iUp = true
			} else {
				j--
				jDown = true
			}
		}
	}
}