package main

import "fmt"

func main() {

	fmt.Println("Hello, 这里是盛最多水的容器!")

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea := maxArea(height)
	fmt.Println(maxArea)
}

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	//遍历指针，然后小的一侧往内缩
	for left < right {
		area := 0
		//得出宽
		width := right - left
		//	比对出最矮的一侧，作为高
		if height[left] < height[right] {
			area = width * height[left]
			left++
		} else {
			area = width * height[right]
			right--
		}
		//	判断面积是否最大
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
