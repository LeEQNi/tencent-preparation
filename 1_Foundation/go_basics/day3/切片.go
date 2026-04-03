package main

import "fmt"

func main() {
	// 1. 定义切片
	var s []int
	fmt.Println("空切片:", s)

	// 2. 初始化切片
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("初始化:", nums)

	// 3. 长度和容量
	fmt.Println("len:", len(nums), " cap:", cap(nums))

	// 4. 切片截取 [起始:结束)
	fmt.Println("nums[1:4]:", nums[1:4]) // [2 3 4]

	// 5. 追加元素
	nums = append(nums, 6, 7)
	fmt.Println("append后:", nums)

	// 6. 删除元素（index=2）
	nums = append(nums[:2], nums[3:]...)
	fmt.Println("删除index=2后:", nums)

	// 7. 插入元素（index=2 插入99）
	nums = append(nums[:2], append([]int{99}, nums[2:]...)...)
	fmt.Println("插入后:", nums)

	// 8. 遍历
	for i, v := range nums {
		fmt.Printf("index:%d value:%d\n", i, v)
	}
}
