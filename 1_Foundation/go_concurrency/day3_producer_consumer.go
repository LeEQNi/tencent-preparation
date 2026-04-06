package main

import (
	"fmt"
	"sync" //同步原语：WaitGroup
	"time" //睡眠模拟业务耗时
)

func main() {
	//	创建带缓冲的通道（缓冲大小10）
	//作用：生产者生产十个以内的数据时不会发生阻塞
	ch := make(chan int, 10)
	//等待一组协程goroutine执行完毕
	//var wg sync.WaitGroup

	// 修改：等待生产者
	var producerWg sync.WaitGroup

	//	启动3个生产者协程
	for i := 1; i <= 3; i++ {
		//等待组计数器+1
		producerWg.Add(1)
		//启动生产者
		go producer(i, ch, &producerWg)
	}

	// 增加单独的协程：生产者完成 -> 关闭通道
	go func() {
		producerWg.Wait()
		close(ch)
		fmt.Printf("所有生产者已完成，通道关闭")
	}()

	// 修改：等待消费者
	var consumerWg sync.WaitGroup

	//	启动2个消费者协程
	for i := 1; i <= 2; i++ {
		//等待组计数器+1
		consumerWg.Add(1)
		//启动消费者
		go consumer(i, ch, &consumerWg)
	}

	//	阻塞：等待所有生产者+消费者 Done()
	//wg.Wait()
	//	关闭通道，通知消费者无数据准备退出
	//close(ch)

	//	睡眠等待消费者处理剩余数据（不优雅）
	//time.Sleep(500 * time.Millisecond)

	// 修改：等待所有消费者完成（优雅）
	consumerWg.Wait()
	fmt.Println("所有任务完成")
}

// 生产者：只写通道 chan<-
func producer(id int, ch chan<- int /*生产者通道*/, wg *sync.WaitGroup) {
	// 函数推出，计数器-1
	defer wg.Done()
	//每个生产者生产3个数据
	for i := 1; i <= 3; i++ {
		//生成产品编号
		product := id*100 + i
		//数据写入通道
		ch <- product
		fmt.Printf("生产者%d -> 产品%03d\n", id, product)
		//模拟生产耗时
		time.Sleep(time.Millisecond * 200)
	}
}

// 消费者：只读同道 <-chan
func consumer(id int, ch <-chan int /*消费者通道*/, wg *sync.WaitGroup) {
	// 函数退出，计数器-1
	defer wg.Done()
	// 遍历通道：通道未关闭且有数据就一直读
	// 通道关闭+数据读完：自动退出循环
	for product := range ch {
		fmt.Printf("消费者%d <- 产品%03d\n", id, product)
		//模拟消费耗时
		time.Sleep(time.Millisecond * 300)
	}
}
