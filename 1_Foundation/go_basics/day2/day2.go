package main

import (
	"fmt"
	"math"
	"strings"
)

//练习1： 变量和类型

func basicTypes() {
	var (
		name   string  = "Go学习者"
		age    int     = 21
		score  float64 = 95.5
		isPass bool    = true
	)

	fmt.Printf("姓名: %s , 年龄: %d , 分数: %.1f , 通过: %v\n ", name, age, score, isPass)
}

// 练习2：控制结构
func controlStructures() {
	fmt.Println("\n=== 控制结构练习 ===")

	// if-else，判断
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d是偶数 ", i)
		} else {
			fmt.Printf("%d是奇数 ", i)
		}
	}

	fmt.Println()

	// switch，选择
	grade := 85
	switch {
	case grade >= 90:
		fmt.Println("优秀")
	case grade >= 80:
		fmt.Println("良好")
	case grade >= 70:
		fmt.Println("中等")
	default:
		fmt.Println("需要努力")
	}
}

// 练习3：函数
func calculateCircle(radius float64) (area, perimeter float64) {
	area = math.Pi * radius * radius
	perimeter = 2 * math.Pi * radius
	return
}

// 练习4：字符串处理
func stringOperations() {
	text := "Go语言并发编程实战"
	fmt.Println("\n=== 字符串操作 ===")
	fmt.Println("原始文本:", text)
	fmt.Println("大写:", strings.ToUpper("go"))
	//判断是否存在，存在则输出true
	fmt.Println("包含'并发':", strings.Contains(text, "并发"))
	//用数组分割字符串，并设立基准为‘，’，然后在每个分割字符中加入空格
	fmt.Println("分割:", strings.Split("a,b,c,d", ","))
}
func main() {
	//输出基本类型
	basicTypes()
	//控制转换
	controlStructures()

	// 函数调用
	//输入半径得面积和周长
	area, perimeter := calculateCircle(5.0)
	fmt.Printf("\n圆(半径=5) 面积: %.2f, 周长: %.2f\n", area, perimeter)

	//字符串处理
	stringOperations()
}
