package main

import "fmt"

// 这道题直接公式推导出来了
// f(m,n) = f(m-1)(n) + f(m)(n-1)
// f(2,2) = 2
// f(1,n) = 1
// f(n,1) = 1
// dp
func uniquePaths(m int, n int) int {
	dp := make(map[[2]int]int)
	for i := 1; i <= m; i++ {
		dp[[2]int{i, 1}] = 1
	}
	for i := 1; i <= n; i++ {
		dp[[2]int{1, i}] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[[2]int{i, j}] = dp[[2]int{i - 1, j}] + dp[[2]int{i, j - 1}]
		}
	}
	return dp[[2]int{m, n}]
}

func main() {
	fmt.Println(uniquePaths(2, 2))
	fmt.Println(uniquePaths(2, 3)) // 3
	fmt.Println(uniquePaths(3, 3)) // 6

}
