package day1

func twoSum(nums []int, target int) [2]int {
	var returnlist [2]int
	for i := range nums {
		do_sum_count := len(nums) - i - 1
		for j := do_sum_count; j > 0; j-- {
			value := nums[i] + nums[i+j]
			if value == target {
				returnlist[0] = i
				returnlist[1] = j
			}
		}

	}
	return returnlist
}
