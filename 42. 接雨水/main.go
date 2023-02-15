package main

import "fmt"

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 双指针
// tirck: 由于l是已知的最小，所以，移动l， 右边主要满足够大就行
// 出水只需要考虑最小
// area[i] = min(l,r) - h[i]
// l,r  if r < l, r--; else l++
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	lmax, rmax := height[l], height[r]
	area := 0
	for l < r {
		if lmax <= rmax {
			l = l + 1
			lmax = Max(height[l], lmax)
			// min l,
			// r = lmax < rmax , lmax
			// lmax < height[i], 0
			area = area + lmax - height[l]
		} else {
			r = r - 1
			rmax = Max(height[r], rmax)
			area = area + rmax - height[r]
		}
	}
	return area
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
