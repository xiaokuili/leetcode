package main

import "fmt"

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：每次只能向下或者向右移动一步。

//

// 示例 1：

// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 示例 2：

// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
//

// 提示：

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
// 通过次数455,777提交次数656,096

// 构建dp

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	dp := make(map[[2]int]int)
	m, n := len(grid), len(grid[0])

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			key := [2]int{i, j}

			if i == 1 {
				sum := 0
				for t := 0; t < j; t++ {
					sum += grid[0][t]
				}
				dp[key] = sum
				continue
			}
			if j == 1 {
				sum := 0
				for t := 0; t < i; t++ {
					sum += grid[t][0]
				}
				dp[key] = sum
				continue
			}
			dp[key] = Min(dp[[2]int{i - 1, j}], dp[[2]int{i, j - 1}]) + grid[i-1][j-1]
			fmt.Println(dp)
		}
	}
	return dp[[2]int{m, n}]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 7
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))

}
