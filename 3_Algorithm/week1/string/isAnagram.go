package main

import (
	"fmt"
)

func main() {

	fmt.Println("这里是有效的字母异位词！")

	s := "anagram"

	t := "nagaram"

	bo := isAnagram(s, t)
	fmt.Println(bo)
}

func isAnagram(s string, t string) bool {
	// 设定一个map，存储s，并且给每个字符添加一个顺序，然后依次遍历t中字符
	// 如果字符完全一样，并且顺序不完全相同，则返回true
	// 时间复杂度：o（n^2）
	//错误题解：
	/*flag := false
	// 词计数，确保不相同，一开始为0
	charCorrect, charNum := 0, 0
	// map存储s中的字符，并标记顺序
	charMap := make(map[int32]int)
	for i, chars := range s {
		charMap[chars] = i
		for j, chart := range t {
			if chart == chars {
				charCorrect++
				if i == j {
					charNum++
				}
			}
		}
	}
	if charNum != len(s) && charCorrect == len(s) {
		flag = true
	}
	return flag
	*/

	//	方法1：排序 时间复杂度：o（nlogn)
	//	思路：用go语言中特有的排序方法来对比两个字符串
	//	首先，如果完全一样，直接返回false，不符合要求
	/*
		if s == t {
			return false
		}
		//	用sort方法来排序比较：
		//	先将两个字符串转换成字符数组
		s1, s2 := []byte(s), []byte(t)
		//	开始排序
		sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
		sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
		//	返回比较结果
		return string(s1) == string(s2)
	*/

	//	方法2：哈希表 时间复杂度：o（n）
	//	思路：这里的关键就是：t是s的变位词等价于[两个字符串不同等且两个字符串中出现的字符的种类和频次相同]
	/*if s == t {
		return false
	}
	//	由于题目是只有26个小写字母，所以这里是[26]int作为种类范围
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
	*/
	// len（s）就是s的字符长度，如果不一样就说明里面的字符不同
	// s和t相同就说明完全一样，也不用测了
	if len(s) != len(t) || s == t {
		return false
	}
	cnt := make(map[rune]int)

	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}

	return true
}
