package main

import (
	"fmt"
	"time"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	done := make(chan struct{})
	go produceEven(even)
	go produceOdd(odd)
	go timeout(done)
	for {
		select {
		case e := <-even:
			fmt.Printf("Even: %d\n", e)
		case o := <-odd:
			fmt.Printf("Odd: %d\n", o)
		case <-done:
			fmt.Println("Done")
			return
		default:
			fmt.Println("Nothing received")
		}
	}
}

func produceEven(nums chan int) {
	for i := 0; ; i += 2 {
		nums <- i
	}
}

func produceOdd(nums chan int) {
	for i := 1; ; i += 2 {
		nums <- i
	}
}

func timeout(done chan struct{}) {
	time.Sleep(2 * time.Second)
	done <- struct{}{}
}
