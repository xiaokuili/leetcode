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

package main

func minWindow(s string, t string) string {
	lt, ls := len(t), len(s)
	pre := make([][]int, 0)
	for i := lt - 1; i >= 0; i-- {
		cur := make([][]int, 0)
		for j := 0; j < ls; j++ {
			if s[j] == t[i] {
				short := 10000000
				m, n := j, j
				for h := 0; h < len(pre); h++ {
					p := pre[h]
					if j < p[0] {
						if p[1]-j < short {
							short = p[1] - j
							m, n = j, p[1]
						}
					}
					if j > p[1] {
						if j-p[0] < short {
							short = j - p[0]
							m, n = p[0], j
						}
					}
					if j > p[0] && j < p[1] {
						if p[1]-p[0] < short {
							short = p[1] - p[0]
							m, n = p[0], p[1]
						}
					}
				}
				cur = append(cur, []int{m, n})
			}
		}
		pre = cur
	}
	short := 1000000
	m, n := 0, 0
	for i := 0; i < len(pre); i++ {
		p := pre[i]
		if p[1]-p[0] < short {
			m, n = p[0], p[1]
		}
	}
	return s[m:n]
}

func main() {

}
