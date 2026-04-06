# 攻坚计划 - 启动周

## 本周目标（2026.4.1-4.6）
1. [x] 完成开发环境的搭建
2. [x] 掌握Git基础操作
3. [x] 学习Markdown文档编写
4. [ ] 完成操作系统第一章学习
5. [ ] 注册leetcode网站，并完成以go语言为基础的十五道算法题目
6. [ ] 注册掘金网站用户，完成两篇学习记录和科研分享文章，每篇不低于一千五百字
7. [ ] 完成一个web代码的实现

## [x] 今日任务（2026.4.1）day1
### 上午
- [x] 创建项目文件夹结构
- [x] 学习Markdown语法
- [x] 安装Git并配置

### 下午
- [go] 选择主语言（Go/Java）并搭建环境
    1. [x] 下载Go 1.26.1安装包
    2. [x] 配置环境变量
    3. [x] 创建测试程序
        1. [x] 在01_function下创建了day1.go中写了一个程序
        2. [x] 记录在04_notes中

### 晚上
- [x] 阅读操作系统第一章（从《OSTEP》开始）

### 遇到的问题
1. **问题描述**：没找到中文版的操作系统书籍，并且打游戏打了很久
2. **解决方案**：试着去问问ai了,但今天已经断电，只能就此作罢。减少游戏时间，把重心放在练习技术和学习知识上来  

### 明日计划

- 开始操作系统学习
- 学习Go语言的基础
- 完成算法练习
  
## [] 今日任务（2026.4.2）day2

### 上午
- [x] 完成Go语言基础语法的学习
- [x] 完成以下代码块的练习：

```go
// 1_Foundation/go_basics/day2.go
package main

import (
    "fmt"
    "math"
    "strings"
)

// 练习1：变量和类型
func basicTypes() {
    var (
        name   string = "Go学习者"
        age    int    = 21
        score  float64 = 95.5
        isPass bool   = true
    )
    
    fmt.Printf("姓名: %s, 年龄: %d, 分数: %.1f, 通过: %v\n", 
        name, age, score, isPass)
}

// 练习2：控制结构
func controlStructures() {
    fmt.Println("\n=== 控制结构练习 ===")
    
    // if-else
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Printf("%d是偶数 ", i)
        } else {
            fmt.Printf("%d是奇数 ", i)
        }
    }
    
    fmt.Println()
    
    // switch
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
    fmt.Println("原始:", text)
    fmt.Println("大写:", strings.ToUpper("go"))
    fmt.Println("包含'并发':", strings.Contains(text, "并发"))
    fmt.Println("分割:", strings.Split("a,b,c,d", ","))
}

func main() {
    basicTypes()
    controlStructures()
    
    // 函数调用
    area, perimeter := calculateCircle(5.0)
    fmt.Printf("\n圆(半径=5) 面积: %.2f, 周长: %.2f\n", area, perimeter)
    
    stringOperations()
}
```

### 下午
- [X] 学习结构体（struct）、切片（slice）、集合（map）
- [x] 完成数据结构练习：
```go
// 1_Foundation/go_basics/day2_data_structures.go
package main

import (
    "fmt"
    "sort"
)

// 用户结构体
type User struct {
    ID       int
    Username string
    Email    string
    Scores   []int
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

func main() {
    dataStructurePractice()
}
```

### 晚上
- [x] leetCode: 完成[数组] 类别的3道题，并记录在3_Algorithm中
    1. [X] 两数之和
    2. [x] 移动零（283）
    3. [x] 盛最多水的容器（11）
    4. [x] 使用bash来创建目录：
```bash
cd ~/Desktop/Tencent_2025/3_Algorithm
mkdir -p week1/array
cd week1/array
```
- [] 阅读第一章的操作系统
- [X] 在4_Notes 中记录学习心得
- [X] git提交今日代码，和学习记录
  

### 遇到的问题
1. **问题描述**：不知道怎么向md中添加代码块
   **解决方案**：询问豆包后发现可以像43行那样添加语法，注意点是不要缩进，空格，以及上下两行要为空,并且要以122行那样结尾

2. **问题描述**：运行代码的时候发现vs code不支持输入内容，因为这是只读编译器
   **解决方案**：使用idea并快速配置了环境‘】j

3. **问题描述**：跨包调用的时候，发现go项目结构不对
   **解决方案**：询问Qoder后发现，必须要在总项目框架下生成，检查go.mod是否正常生成

4. **问题描述**：晚上花了三个小时打游戏，还有一些时候效率太低了
   **解决方案**：决定改变自己的作息
### 明日计划

- 继续学习操作系统第二章
- 学完剩下的go语言基础
- 尝试制作一个项目框架，进行web项目搭建
- 阅读一篇技术论文

## [x] 今日任务(2026.4.3)day3

    由于今日也没完成任务，所以安排完成昨天的任务，比如说数据结构的学习，但是算法什么的依旧没完成

### 下午
- [x] 完成slice学习
  
### 晚上
- [x] 完成map学习
- [X] 在4_Notes 中记录学习心得

### 遇到的问题
1. **问题描述**：时间安排紊乱，计划安排严重拖沓
   **解决方案**：安排一个松弛时间，每周至少安排半天来缓解冲击

## [] 今日任务(2026.4.5)day5

**今天的核心任务就是补全前面几天欠缺的，同时适当推进剩下的任务**

### 上午
- [x] leetCode: 完成[数组] 类别的3道题，并记录在3_Algorithm中
    1. [x] 两数之和
    2. [x] 移动零（283）
    3. [x] 盛最多水的容器（11）
### 下午
- [] 生产者-消费者模型（1.5h）
- [] 简化版HTTP服务（1.5h）
- [] 确保两个任务能正常运行测试
### 晚上
- [] 完成[字符串]内容：
  1. [x] 反转字符串（344） 
  2. [] 有效的字母异位词（
- [] 发布第一篇技术文章：
  
### 遇到的问题
1. **问题描述**：
   **解决方案**：


### 明日安排
目标：完成Day4+补全所有算法
上午：数据库连接，redis缓存
下午：项目结构搭建，补全算法
晚上：发布第二篇文章总结本周学习，做下周规划