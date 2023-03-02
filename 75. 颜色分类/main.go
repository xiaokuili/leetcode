package main

import "fmt"

// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

// 必须在不使用库内置的 sort 函数的情况下解决这个问题。

//

// 示例 1：

// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]
// 示例 2：

// 输入：nums = [2,0,1]
// 输出：[0,1,2]
//

// 提示：

// n == nums.length
// 1 <= n <= 300
// nums[i] 为 0、1 或 2
//

// 进阶：

// 你能想出一个仅使用常数空间的一趟扫描算法吗？
// 通过次数505,189提交次数836,884

// 经过两次循环，
// 第一次排前部分
// 第二次排后半部分

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] > nums[r] {
			tmp := nums[r]
			nums[r] = nums[l]
			nums[l] = tmp
		}
		if nums[l] == nums[r] && nums[l] == 1 {
			r--
		}

		if nums[l] == 0 {
			l++
		}
		if nums[r] == 2 {
			r--
		}

	}
	r = len(nums) - 1
	for l < r {
		if nums[r] == 2 {
			r = r - 1
		}
		if nums[l] > nums[r] {
			tmp := nums[r]
			nums[r] = nums[l]
			nums[l] = tmp
		}
		if nums[l] == nums[r] && nums[l] == 1 {
			l++
		}
		if nums[l] == 0 {
			l++
		}
		if nums[r] == 2 {
			r--
		}
	}
}

func main() {
	var nums []int
	nums = []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{0, 1}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{0, 2, 2, 2, 0, 2, 1, 1}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{1, 0, 1, 2, 1, 0, 0, 0}
	sortColors(nums)
	fmt.Println(nums)

}
