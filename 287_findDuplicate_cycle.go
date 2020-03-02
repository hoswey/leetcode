func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[0]

	for  {
		slow = nums[slow]
		fast = nums[nums[fast]]
        if slow == fast {
            break
        }
	}

    fmt.Printf("1\n")
    
	p := nums[0]
	for p != slow {
		p = nums[p]
		slow = nums[slow]
	}
    fmt.Printf("2\n")

	return p
}