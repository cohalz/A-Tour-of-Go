package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3
	//fatal error: all goroutines are asleep - deadlock!
	//goroutine 1 [chan send]:

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//=> all goroutines are asleep - deadlock!
	//goroutine 1 [chan receive]:
}
