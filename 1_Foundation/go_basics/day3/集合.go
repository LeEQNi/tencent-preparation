package main

import "fmt"

func main() {
	// 1. 【创建map】3种方式（最常用：make / 直接初始化）
	// 方式1：make创建空map（key string，value int）
	var m1 map[string]int
	m1 = make(map[string]int) // 必须初始化才能用，否则nil无法赋值

	// 方式2：简写（推荐）
	m2 := make(map[string]string)

	// 方式3：直接初始化（带初始值）
	m3 := map[string]int{
		"苹果": 5,
		"香蕉": 3,
	}

	// 2. 【赋值/修改】key不存在=新增，存在=修改
	m1["张三"] = 18
	m1["李四"] = 20
	m1["张三"] = 19 // 覆盖原有值

	m2["name"] = "Go学习"
	m2["version"] = "1.22"

	// 3. 【取值】两种方式
	// 方式1：直接取（key不存在，返回value类型零值）
	age := m1["张三"]
	fmt.Println("张三年龄：", age) // 19

	// 方式2：带判断（推荐，防止key不存在）
	value, ok := m1["王五"]
	if ok {
		fmt.Println("王五年龄：", value)
	} else {
		fmt.Println("王五不存在") // 输出这个
	}

	// 4. 【遍历】for range（无序！Go map不保证顺序）
	fmt.Println("\n--- 遍历m3 ---")
	for k, v := range m3 {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	// 只遍历key
	for k := range m3 {
		fmt.Println("key：", k)
	}

	// 5. 【删除元素】delete(map, key)
	delete(m1, "李四")
	fmt.Println("\n删除李四后：", m1) // map[张三:19]

	// 6. 【长度】len(map) → 元素个数
	fmt.Println("m3长度：", len(m3)) // 2

	// 7. 【清空map】无直接清空方法，重新make即可
	m3 = make(map[string]int)
	fmt.Println("清空后m3：", m3) // map[]

	// 8. 【嵌套map】value是map（常用场景：复杂数据）
	student := make(map[string]map[string]int)
	student["小明"] = map[string]int{"语文": 90, "数学": 95}
	fmt.Println("\n小明语文成绩：", student["小明"]["语文"]) // 90
}
