// 1_Foundation/go_basics/day2_data_structures.go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// 定义一个用户结构体
type User struct {
	ID       int
	Username string
	Email    string
	Scores   []int
}

// 结构体方法
func (u User) GetHighestScore() int {
	if len(u.Scores) == 0 {
		return 0
	}
	max := u.Scores[0]
	for _, score := range u.Scores {
		if score > max {
			max = score
		}
	}
	return max
}

// 补全：更新用户信息的方法
func (u *User) UpdateEmail(newEmail string) error {
	if !strings.Contains(newEmail, "@") {
		return fmt.Errorf("邮箱格式错误")
	}
	u.Email = newEmail
	return nil
}

func dataStructurePractice() {
	fmt.Println("=== 数据结构练习 ===")

	// 1. 切片
	fmt.Println("\n1. 切片操作:")
	scores := []int{85, 90, 78, 92, 88}
	fmt.Println("原始切片:", scores)

	// 添加元素
	scores = append(scores, 95)
	fmt.Println("添加后:", scores)

	// 排序
	sort.Ints(scores)
	fmt.Println("排序后:", scores)

	// 2. 映射
	fmt.Println("\n2. 映射操作:")
	userMap := map[string]User{
		"alice": {1, "alice", "alice@example.com", []int{85, 90}},
		"bob":   {2, "bob", "bob@example.com", []int{78, 92}},
	}

	// 遍历map
	for username, user := range userMap {
		fmt.Printf("%s: %s (%s)\n", username, user.Username, user.Email)
	}

	// 3. 结构体方法
	fmt.Println("\n3. 结构体方法:")
	user := User{
		ID:       1001,
		Username: "tencent_learner",
		Email:    "learn@tencent.com",
		Scores:   []int{90, 85, 95},
	}

	// 计算平均分
	avg := user.AverageScore()
	fmt.Printf("用户 %s 平均分: %.1f\n", user.Username, avg)

	// 4. 错误处理
	fmt.Println("\n4. 错误处理:")
	if user, err := FindUserByID(userMap, 1); err == nil {
		fmt.Println("找到用户:", user.Username)
	} else {
		fmt.Println("错误:", err)
	}
}

// 结构体方法
func (u User) AverageScore() float64 {
	if len(u.Scores) == 0 {
		return 0.0
	}

	total := 0
	for _, score := range u.Scores {
		total += score
	}
	return float64(total) / float64(len(u.Scores))
}

// 错误处理示例
func FindUserByID(users map[string]User, id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("用户ID %d 不存在", id)
}

// 启动main方法
func main() {
	//调用数据结构体练习方法
	dataStructurePractice()
}
