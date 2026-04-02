package main

// 调用一个输出包
// import "fmt"

// 当要调用的包比较多，我们可以这样

import (
	"fmt"
	// 当然不调用也可以，但这里我们先调用，同时一定要锁定根目录的文件,以下是示范，请你使用的时候换成自己的vertion包的路径地址
	"day2/vertion"
)

// 现在定义全局变量
// 即使没人调用也不会报红
var a = 12

// 全局变量中定义多个变量
var (
	s1 string = "在外面"
	s2 string = "不报错"
)

func hello() {
	// 在这个方法中无法调用main中变量
	fmt.Println("你好")

	// 但在这个方法中可以用a，因为a是全局变量
	fmt.Println(a)
}

func main() {

	// 这样就能调用另一个vertion包中的内容了
	fmt.Println(vertion.Vertion)
	// 但不能调用小写的vertion,这样会报错
	// fmt.Println(vertion)

	// 现在定义局部变量

	// 先声明
	var name string
	// 再赋值
	name = "ad"
	// 定义了不使用name就会标红
	fmt.Println(name)

	// 直接声明赋值
	var name1 string = "Lee"
	fmt.Println(name1)

	// 不用定义类型，也能自动识别
	var name2 = "lili"
	fmt.Println(name2)

	// 还有快速定义的方法，这也叫短声明
	name3 := "niuniu"
	fmt.Println(name3)

	//调用方法
	hello()

	// 定义多个变量
	var a1, a2 = 2, 3
	fmt.Println(a1, a2)

}
