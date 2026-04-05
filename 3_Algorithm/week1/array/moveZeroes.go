package main

import "fmt"

func main() {

	fmt.Println("Hello, 我是移动零！")
	nums := []int{0, 1, 0, 3, 12}
	//moveZerose函数声明没有返回值，但是以下调用却接收了返回值
	//nums1 := []int{}
	//nums1 = moveZeroes(nums)

	//Go的切片是引用传递，函数里修改会直接改原数组，不需要额外赋值
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	// 双指针方法，设定两个一左一右指针
	// 1.左指针指向**已经处理好的序列的尾部**，所以左指针的左边都是非零数
	// 2.右指针的左边到左指针都是0
	// 每次交换都是将左指针的零与右指针的非零数交换，且非零数的相对顺序没变
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
