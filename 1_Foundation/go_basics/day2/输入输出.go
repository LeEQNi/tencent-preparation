package main

import "fmt"

func main() {

	// 首先是输出

	// 三种print，print非格式化输出不换行不空格；println自动换行；printf是格式化输出字符串
	// 不声明类型不能输出整数等类型，默认是字符串
	fmt.Print(1)
	fmt.Printf("1\n")
	fmt.Printf("%s是个超绝大美女", "灰灰大王")
	fmt.Println("我没有换行，我在美女后面")
	fmt.Print("我换行了，我在美女下面")

	// 格式化输出的一些常用符号
	// 可以作为任何值的占位符输出
	fmt.Printf("%v\n", "你好")
	// 打印类型T
	fmt.Printf("%v %T\n", "niu", "er")
	// 整数
	fmt.Printf("%d\n", 3)
	// 小数
	fmt.Printf("%.2f\n", 1.25)
	// 字符串
	fmt.Printf("%s\n", "length")
	// 用go的语法格式输出，很适合打印空字符串
	fmt.Printf("%#v\n", "")

	// 还有一个常用的将格式化后的内容赋值给一个变量
	name := fmt.Sprintf("%v", "你好")
	fmt.Println(name)

	// 输入

	fmt.Print("请输入你的名字:")
	var name1 string
	// &指针,指向这个定义的变量,我个人感觉变量存在方法里,应该是在栈中存储
	fmt.Scan(&name1)
	fmt.Println(name1)
	fmt.Print("请输入你的年龄:")
	var age int
	// 这里额外展示一次表示不限数据类型
	// fmt.Scan(&age)
	// 可以用另一个短定义的方式接住
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}
