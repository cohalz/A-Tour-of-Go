package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	/*
		1個目のバッファ
				c <- 0
						x, y = 1, 1

								2個目のバッファ
										c <- 1
												x, y = 1, 2
														の繰り返し
	*/
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	c := make(chan int, 10)
	//go fibonacci(10, c)
	go fibonacci(cap(c), c)

	//for i := range c は、チャネルが閉じられるまで、チャネルから値を繰り返し受信し続けます。
	//closeがないと11回目も送信しようとしてfatal error: all goroutines are asleep - deadlock!
	for i := range c {
		fmt.Println(i)
	}
}
