func drop(j int, nums []int) []int {
	n := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if j != i {
			n = append(n, nums[i])
		}
	}
	return n
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		terms := permute(drop(i, nums))
		for j := 0; j < len(terms); j++ {
			terms[j] = append([]int{cur}, terms[j]...)

			result = append(result, terms[j])
		}
	}
	return result

}