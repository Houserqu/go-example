package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	demo4()
}

func demo1() {
	go func() {
		fmt.Println("1")
	}()
	go func() {
		fmt.Println("2")
	}()
	go func() {
		fmt.Println("3")
	}()
	go func() {
		fmt.Println("4")
	}()

	time.Sleep(2 * 1e9)
}

func demo2() {
	runtime.GOMAXPROCS(1)  // 指定最大 P 为 1，从而管理协程最多的线程为 1 个
	wg := sync.WaitGroup{} // 控制等待所有协程都执行完再退出程序
	wg.Add(2)
	// 运行一个协程
	go func() {
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
		wg.Done()
	}()

	// 运行第二个协程
	go func() {
		fmt.Println(65)
		fmt.Println(66)
		// 设置个睡眠，让该协程执行超时而被挂起，引起超时调度
		time.Sleep(time.Second)
		fmt.Println(67)
		wg.Done()
	}()
	wg.Wait()
}

// channel
func demo3() {
	ch1 := make(chan string)

	go func() {
		ch1 <- "1"
		ch1 <- "2"
		ch1 <- "3"
	}()

	go func() {
		for {
			in := <-ch1
			fmt.Println(in)
		}
	}()

	time.Sleep(1e9)
}

func demo4() {
	ch := make(chan string)
	go func() {
		ch <- "1111"
		ch <- "2222"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
