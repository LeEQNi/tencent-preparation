package main

import "fmt"

func main() {

	fmt.Println("这里是反转字符串！")
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	//不能这么直接输出
	//fmt.Println(s) //[111 108 108 101 104]

	//	正确输出格式
	fmt.Print("[")
	for i, ch := range s {
		fmt.Printf("%q", string(ch))
		//	不是最后一个字符就加逗号
		if i < len(s)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}

func reverseString(s []byte) {
	//	设定两个指针，头尾交换
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
