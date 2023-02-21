package main

import (
	"fmt"
)

// 判断出现次数
// 26个字母出现次数，如果出现次数形同则相等，则添加到结果中
// 映射用的好呀

func groupAnagrams(strs []string) [][]string {
	result := make(map[[26]int][]string, 0)

	for i := 0; i < len(strs); i++ {
		var keys [26]int

		s := strs[i]
		for j := 0; j < len(s); j++ {
			c := s[j]
			keys[c-'a'] = keys[c-'a'] + 1

		}
		fmt.Println(keys)

		if value, ok := result[keys]; ok {

			value = append(value, s)
			result[keys] = value
		} else {
			result[keys] = []string{s}
		}

	}
	res := make([][]string, 0)
	for _, v := range result {
		res = append(res, v)
	}
	return res
}

func main() {
	// 	输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
	// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
	// 示例 2:

	// 输入: strs = [""]
	// 输出: [[""]]
	// 示例 3:

	// 输入: strs = ["a"]
	// 输出: [["a"]]

	// result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	// fmt.Println(result) // [["bat"],["nat","tan"],["ate","eat","tea"]]

	// result := groupAnagrams([]string{""})
	// fmt.Println(result) // [[""]]

	// result := groupAnagrams([]string{"a"})
	// fmt.Println(result) // [["a"]]

	result := groupAnagrams([]string{"bdddddddddd", "bbbbbbbbbbc"})
	fmt.Println(result) // [["a"]]

}
