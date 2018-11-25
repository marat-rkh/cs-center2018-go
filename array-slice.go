package main

import "fmt"

func arrays() {
	nums := [3]int{1, 2, 3}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
}

func slices() {
	words := []string{`ready`, `steady`}
	words = append(words, `go!`)
	fmt.Println(words) // Prints: [ready steady go!]

	nums := [3]int{1, 2, 3}
	slice := nums[0:2]
	fmt.Println(slice) // Prints: [1 2]

	slice = append(slice, 4)
	fmt.Println(slice) // Prints: [1 2 4]
	fmt.Println(nums)  // Prints: [1 2 4]

	slice = append(slice, 5)
	fmt.Println(slice) // Prints: [1 2 4 5]
	fmt.Println(nums)  // Prints: [1 2 4]
}
