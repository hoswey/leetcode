func nextPermutation(nums []int)  {

	if len(nums) <= 1 {
		return
	}

	//1、从后往前找出第一对 a[i] > a[i-1]
	for  i := len(nums) - 1; i >= 1; i-- {

		if nums[i] > nums[i-1] {
			//2、从后往前， 找出第一个 a[k] > a[i-1]的， swap

			j := len(nums) - 1;			
			for {
				if nums[j] <= nums[i-1] {
					j--
				}  else {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					// 反转a[i:end]
					reverse(nums[i:])
					return					
				}

			}
		}

	}

	//3、假如未找到一组合法对，表示该排列是最大， 则反转a[i:end]
	reverse(nums)

}

func reverse(arr []int) {

	lo := 0
	hi := len(arr) - 1

	for lo < hi {
		arr[lo], arr[hi] = arr[hi], arr[lo]
		lo++
		hi--
	}

}