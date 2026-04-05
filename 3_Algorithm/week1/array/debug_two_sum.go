package main

import "fmt"

func main() {
	//	这是程序入口
	fmt.Println("Hello, 两数之和！")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("nums: %v, target: %d, result: %v\n", nums, target, result)
}

// 两数之和函数
// 给予一个整数数组，一个目标数，通过两两相加，若是等于目标数，则返回两个下标
func twoSum(nums []int, target int) []int {
	// 暴力双循环 - 时间复杂度 O(n²)
	// 遍历每个数，看它后面有没有能和它加起来等于target的数
	/*
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					//	找到目标下标，返回
					return []int{i, j}//[0,1]
				}
			}
		}
	*/
	// 哈希表 - 时间复杂度0(n)
	// 创建一个map记录每个数需要的“另一半”的下标
	// 创建一个map，key：存值，value：存下标
	numMap := make(map[int]int)
	// 遍历nums数组中的数字,i是当前索引，num是当前值
	//for i:=0;i<len(nums);i++ 这是我的版本
	for i, num := range nums {
		//	找到另一半
		complement := target - num
		//	判断是否能在numMap中找到另一半
		//if idx, found := numMap(complement) found 这是我写的有错误的版本
		if idx, found := numMap[complement]; found {
			//	找到了
			//	idx是存储的“另一半”的下标，i是当前数的下标
			return []int{idx, i}
		}
		//	没找到，把当前数字存入numMap中，这里不会写
		// 这样以后别的数就能找到它
		numMap[num] = i
	}
	//返回一个空数组,没找到
	return []int{} //[]
}
