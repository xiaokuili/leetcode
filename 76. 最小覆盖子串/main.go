// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

//

// 注意：

// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//

// 示例 1：

// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
// 示例 2：

// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
// 示例 3:

// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
// 这道题竟然不是dp
// 遇到是否出现，但是不考虑顺序，就去hash

package main

import "fmt"

func minWindow(s string, t string) string {
	if t == "" || s == "" {
		return ""
	}
	windowT, windowS := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		windowT[t[i]] = windowT[t[i]] + 1
	}
	fmt.Println(windowT)
	need, have := len(windowT), 0
	rstl, rstr := 0, 0
	min := 100000
	l := 0
	for r := 0; r < len(s); r++ {
		windowS[s[r]] = windowS[s[r]] + 1

		if _, ok := windowT[s[r]]; ok && windowS[s[r]] == windowT[s[r]] {
			have++

		}
		for need == have {
			if (r - l + 1) < min {
				min = r - l + 1
				rstl = l
				rstr = r + 1
			}
			windowS[s[l]] = windowS[s[l]] - 1
			if _, ok := windowT[s[l]]; ok && windowS[s[l]] < windowT[s[l]] {
				have = have - 1
			}
			l = l + 1

		}
	}
	return s[rstl:rstr]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // BANC
}
