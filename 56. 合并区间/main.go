package main

import (
	"sort"
)

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

//

// 示例 1：

// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2：

// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 暴力破解
// 合并一个数组和另一个数组

// func merge2Slice(s1, s2 []int) (bool, int, int) {
// 	x, y, x1, y1 := s1[0], s1[1], s2[0], s2[1]
// 	if x1 <= x && x <= y1 {
// 		if y1 < y {
// 			return true, x1, y
// 		}
// 		return true, x1, y1
// 	}
// 	if x1 <= y && y <= y1 {
// 		if x1 < x {
// 			return true, x1, y1
// 		}
// 		return true, x, y1
// 	}
// 	return false, 0, 0
// }

// func merge(intervals [][]int) [][]int {
// 	i := 0
// 	l := len(intervals)
// 	for i < l {
// 		s1 := intervals[i]
// 		for j := 0; j < len(intervals); j++ {
// 			if i != j && intervals[j] != nil {

// 				ok, x, y := merge2Slice(s1, intervals[j])

// 				if ok {
// 					intervals[i] = nil
// 					intervals[j] = []int{x, y}
// 					break
// 				}
// 			}
// 		}
// 		i++

// 	}
// 	result := make([][]int, 0)
// 	for i := 0; i < len(intervals); i++ {
// 		if intervals[i] != nil {
// 			result = append(result, intervals[i])

// 		}
// 	}
// 	return result
// }

// 非暴力破解
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	result := make([][]int, 0)
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		r, inter := result[len(result)-1], intervals[i]
		start := inter[0]
		lastEnd := r[1]
		if start <= lastEnd {
			r[1] = Max(r[1], inter[1])
		} else {
			result = append(result, inter)
		}
	}

	return result
}

func main() {

	// intervals := make([][]int, 0)
	// intervals = append(intervals, []int{1, 3})
	// intervals = append(intervals, []int{2, 6})
	// intervals = append(intervals, []int{8, 10})
	// intervals = append(intervals, []int{15, 18})
	// result := merge(intervals)
	// fmt.Println(result) //[[1,6],[8,10],[15,18]]

	// intervals := make([][]int, 0)
	// intervals = append(intervals, []int{1, 4})
	// intervals = append(intervals, []int{4, 5})
	// result := merge(intervals)
	// fmt.Println(result) //[[1,5]]

	// intervals := make([][]int, 0)
	// intervals = append(intervals, []int{2, 3})
	// intervals = append(intervals, []int{4, 5})
	// intervals = append(intervals, []int{6, 7})
	// intervals = append(intervals, []int{8, 9})
	// intervals = append(intervals, []int{1, 10})

	// result := merge(intervals)
	// fmt.Println(result) //[[1,10]]
}
