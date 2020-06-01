package main

import (
	"sort"
)

//1464.数组中两元素的最大乘积
func maxProduct(nums []int) int {
	max := 0
	for i := 0;i < len(nums);i++{
		for j := i + 1;j < len(nums);j++{
			if (nums[i] - 1) * (nums[j] - 1) > max {
				max = (nums[i] - 1) * (nums[j] - 1)
			}
		}
	}
	return max
}

//1465.切割后面积最大的蛋糕,由于答案可能是一个很大的数字，因此需要将结果对 10^9 + 7 取余后返回
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	height := 0
	length := 0
	horizontalCuts = append(horizontalCuts, 0)
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, 0)
	verticalCuts = append(verticalCuts, w)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	for i := 0;i < len(horizontalCuts) - 1;i++{
		if horizontalCuts[i + 1] - horizontalCuts[i] > height {
			height = horizontalCuts[i + 1] - horizontalCuts[i]
		}
	}
	for j := 0;j < len(verticalCuts) - 1;j++{
		if verticalCuts[j + 1] - verticalCuts[j] > length {
			length = verticalCuts[j + 1] - verticalCuts[j]
		}
	}
	return (height*length) % (1000000000 + 7)
}

//1466. 重新规划路线
func minReorder(n int, connections [][]int) int {

}

