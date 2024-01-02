package main

import "fmt"

func countDown(number int) {
	if number <= 0 {
		return
	}
	fmt.Println("count down starts in: ", number)
	countDown(number - 1)
}

func main() {
	countDown(10)
}
