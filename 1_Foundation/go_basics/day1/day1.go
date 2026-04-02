package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 基础信息
	fmt.Println("=== Go环境信息 ===")
	fmt.Printf("Go版本: %s\n", runtime.Version())
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Printf("CPU核心数: %d\n", runtime.NumCPU())

	// 学习计时开始
	startTime := time.Now()
	fmt.Printf("\n学习开始时间: %v\n", startTime.Format(("2006-01-02 15:04:05")))
	// 第一个Go特性：并发实例
	fmt.Println("\n=== Go并发初体验 ===")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d 执行\n", id)
		}(i)
	}
	// 等待goroutine执行
	time.Sleep(100 * time.Millisecond)
	// 学习时长
	duration := time.Since(startTime)
	fmt.Printf("\n√ 环境搭建完成!耗时:%.2f秒\n", duration.Seconds())
	fmt.Println("明日任务:变量，函数，控制结构")
}
