package main

import "fmt"

// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符
//

// 示例 1：

// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2：

// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// dp
// 分为已经构建和未构建部分
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l2+1)
	for i := 0; i < l2+1; i++ {
		tmp := make([]int, l1+1)
		dp[i] = tmp
	}
	for i := 0; i < len(dp[l2]); i++ {
		dp[l2][i] = len(word1[i:])
	}

	for i := 0; i < len(dp); i++ {
		dp[i][l1] = len(word2[i:])
	}

	for j := len(dp) - 2; j >= 0; j-- {
		for i := len(dp[j]) - 2; i >= 0; i-- {
			if word1[i] == word2[j] {
				dp[j][i] = dp[j+1][i+1]
			} else {
				dp[j][i] = Min(dp[j][i+1], Min(dp[j+1][i+1], dp[j+1][i])) + 1
			}
			// fmt.Println(i, j, dp, dp[j][i], string(word1[i]), string(word2[j]))

		}
	}
	return dp[0][0]

}

func main() {
	// fmt.Println(minDistance("abc", "adc")) // 1
	fmt.Println(minDistance("horse", "ros")) // 3

	fmt.Println(minDistance("intention", "execution")) // 5

}
