package main

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 子数组 是数组中的一个连续部分。

//

// 示例 1：

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 示例 2：

// 输入：nums = [1]
// 输出：1
// 示例 3：

// 输入：nums = [5,4,-1,7,8]
// 输出：23
//

// 提示：

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104

// 左右指针
// 先构建左边，负数直接右滑， 正数停止，如果后面负数小于正数， 构建成功
// 右边，负数停止，遍历，如果后面的可以填充负数，后移，直到不能填充
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	cur := 0
	for i := 0; i < len(nums); i++ {
		if cur < 0 {
			cur = 0
		}
		cur = cur + nums[i]
		maxSum = Max(maxSum, cur)
	}
	return maxSum
}
