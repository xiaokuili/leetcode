package main

import "fmt"

// 如果不是从指针出发，那就可能出错了
// l,r,t,b
// l<=i<r, 开始遍历，每个元素循环四次
// 位置分别为：[b, l + i]， [b +i, r], [t, r -i], [t-i,l]
// 替换：tmp := m[b][l+i], m[b][l+1],m[t-i][l],m[t][r-i],m[b+i][r] = m[t-i][l], m[t][r-i],m[b+i][r],tmp
// 随后l和r缩减

func rotate(matrix [][]int) {
	l, r, t, b := 0, len(matrix)-1, len(matrix)-1, 0
	for l < r {
		fmt.Println(l, r, b, t)
		for i := 0; i < r-l; i++ {
			tmp := matrix[b][l+i]
			fmt.Println(b, i, l+i, tmp, matrix)
			matrix[b][l+i], matrix[t-i][l], matrix[t][r-i], matrix[b+i][r] = matrix[t-i][l], matrix[t][r-i], matrix[b+i][r], tmp
			fmt.Println("-----", matrix)
		}
		l, r, t, b = l+1, r-1, t-1, b+1

	}
}

func main() {
	// matrix := make([][]int, 0)
	// matrix = append(matrix, []int{1, 2, 3})
	// matrix = append(matrix, []int{4, 5, 6})
	// matrix = append(matrix, []int{7, 8, 9})

	// rotate(matrix)
	// fmt.Println(matrix) // [[7 4 1] [8 5 2] [9 6 3]]

	// matrix := make([][]int, 0)
	// matrix = append(matrix, []int{5, 1, 9, 11})
	// matrix = append(matrix, []int{2, 4, 8, 10})
	// matrix = append(matrix, []int{13, 3, 6, 7})
	// matrix = append(matrix, []int{15, 14, 12, 16})

	// rotate(matrix)
	// fmt.Println(matrix) // [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

	matrix := make([][]int, 0)
	matrix = append(matrix, []int{2, 29, 20, 26, 16, 28})
	matrix = append(matrix, []int{12, 27, 9, 25, 13, 21})
	matrix = append(matrix, []int{32, 33, 32, 2, 28, 14})
	matrix = append(matrix, []int{13, 14, 32, 27, 22, 26})
	matrix = append(matrix, []int{33, 1, 20, 7, 21, 7})
	matrix = append(matrix, []int{4, 24, 1, 6, 32, 34})

	rotate(matrix)
	fmt.Println(matrix) // [[4,33,13,32,12,2],[24,1,14,33,27,29],[1,20,32,32,9,20],[6,7,27,2,25,26],[32,21,22,28,13,16],[34,7,26,14,21,28]]

}
