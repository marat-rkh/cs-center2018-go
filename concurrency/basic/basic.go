package basic

import "fmt"

func one() {
	nums := make(chan int) // unbuffered
	go produceOne(nums)
	consumeOne(nums)
}

func produceOne(nums chan int) {
	nums <- 1
}

func consumeOne(nums chan int) {
	n := <-nums
	fmt.Println(n)
}

func many() {
	nums := make(chan int, 5) // buffered
	go produce(nums)
	consume(nums)
}

func produce(nums chan int) {
	for i := 0; i < 10; i++ {
		nums <- i
	}
	close(nums)
}

func consume(nums chan int) {
	for n := range nums {
		fmt.Println(n)
	}
}
