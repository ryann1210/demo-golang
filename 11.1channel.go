package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c \n", id, n)
	//}
	for n := range c {
		fmt.Printf("Worker %d received %d \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	//var c chan int // c == nil 所以不可以用
	var channels [10]chan<- int
	for i := 0; i < len(channels); i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < len(channels); i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < len(channels); i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	bufferedChannel()
	//channelClose()
}

// 发送放通知接收方 数据已经发送完毕
func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	// 缓冲区大小为3 发第四个的时候才会deadlock
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 'a'
	time.Sleep(time.Millisecond)
}
