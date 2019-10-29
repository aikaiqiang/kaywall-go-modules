package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched() 表示让 CPU 把时间片让给别人,下次某个时候继续恢复执行该 goroutine
		runtime.Gosched()
		fmt.Println(s)
	}
}

// 1. channels
// channel 接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock
// 所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出
// 记住应该在生产者的地方关闭 channel，而不是消费的地方去关闭它，这样容易引起 panic
// channel 不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束 range 循环之类的
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func channelTest() {

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	go sum(a[0:3], c)
	z := <-c
	fmt.Println(z)
}

// 2. Buffered Channels
func bufferedChannelTest() {
	c := make(chan int, 1) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

// 3. Range 和 Close
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeTest() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// 判断 channel 是否关闭
	v, ok := <-c
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("channel is closed")
	}
}

// 4. 多个channel， Select：通过 select 可以监听 channel 上的数据流动
func fibonacciUp(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit: // 当 quit 有数据流进来表示要退出循环
			fmt.Println("quit")
			return
		}
	}
}

func multipleChannelTest() {
	c := make(chan int)
	quit := make(chan int)

	// （消费者）开启一个写数据线程
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// （生产者）开启一个写数据线程
	fibonacciUp(c, quit)

}

// timeout
func timeoutTest() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	result := <-o
	println("result = ", result)
}

func main() {
	//go say("world") //开一个新的Goroutines执行
	//say("hello") //当前Goroutines执行

	//channelTest()
	//bufferedChannelTest()
	//rangeTest()
	multipleChannelTest()
	timeoutTest()

	cpu := runtime.NumCPU()
	goroutine := runtime.NumGoroutine()
	gomaxprocs := runtime.GOMAXPROCS(4)

	fmt.Printf("cpu = [%v], goroutineCount = [%v], gomaxprocs = [%v]", cpu, goroutine, gomaxprocs)
}
