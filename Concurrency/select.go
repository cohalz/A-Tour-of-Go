package main

import "fmt"
import "strconv"

//goroutineは送受信以外は不定

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("fib a")
		select {
		case c <- x:
			fmt.Println("fib x: " + strconv.FormatInt(int64(x), 10))
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//この前にgoroutineがmainしかないためダメ
	//fibonacci(c, quit)

	go func() {
		for i := 0; i < 10; i++ {
			i64 := int64(i)
			fmt.Println("func i: " + strconv.FormatInt(i64, 10) + " start")
			fmt.Print("func ")
			fmt.Print(<-c)
			fmt.Println(" <- c")
			fmt.Println("func i: " + strconv.FormatInt(i64, 10) + " end\n")
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

/*


	{fibonacciを実行}
	fib aを表示
	受信の準備が出来ていないので，selectで止まる
	{fib止まる 受信準備待ち}
	[func実行]
	func i: 0 start
	func
	[<- c に来たので受信準備完了を知らせ止まる 受信来たら再開]
	{fib動く}
	c <- x
	[funcが受信，再開]
	{fib止まる}
	0 <- c
	func i: 0 end

	func i: 1 start

	func fib
	[func受信待ち]
	{fib再開}
	x: 0
	fib a
	c <- 1
	<fib止まらない>
	fib x: 1
	fib a
	{fib受信準備待ちで止まる}
	[func受信]
	1 <- c
	func i: 1 end

	func i: 2 start
	func
	[func受信準備完了]
	{fib再開}
	c <- 2
	[funcが受信，再開]
	{fib止まる}
	1 <- c
	func i: 2 end

	の繰り返し

	順序は不定
*/
