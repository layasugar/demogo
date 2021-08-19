// 测试程序终止信号，程序关闭chan保证不丢失数据

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var TestSignalProducerChan = make(chan int)
var TestSignalConsumerChan = make(chan int)

func mainSignal() {
	Signal()

	// 消费
	go func() {
		for consumer := range TestSignalConsumerChan {
			time.Sleep(time.Second)
			fmt.Printf("consumer： %d\r\n", consumer)
		}
	}()

	// 分发
	go func() {
		for producer := range TestSignalProducerChan {
			time.Sleep(time.Second)
			fmt.Printf("producer： %d\r\n", producer)
			TestSignalConsumerChan <- producer
		}
	}()

	// 生产
	for i := 0; i < 10000; i++ {
		time.Sleep(100 * time.Millisecond)
		TestSignalProducerChan <- i
	}
}

func Signal() {
	// 合建chan
	c := make(chan os.Signal)

	// 监听指定信号 ctrl+c kill
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// 信号处理
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGTERM:
				// 关闭consumer
				fmt.Print("SIGTERM 退出\r\n")
				exitFunc()
			case os.Interrupt:
				fmt.Print("user ctrl c\r\n")
				exitFunc()
			default:
				fmt.Println("other", s)
			}
		}
	}()
}

func exitFunc() {
	go func() {
		for i := 0; i < 60; i++ {
			fmt.Printf("程序退出第%d秒\r\n", i)
			time.Sleep(time.Second)
		}
	}()

	close(TestSignalProducerChan)

	time.Sleep(28 * time.Second)

	close(TestSignalConsumerChan)
}
