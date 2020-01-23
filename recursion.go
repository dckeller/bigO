package main 

import "fmt"

func countdown(i int) {
	fmt.Println(i)

	if i <= 0 {  // Base case when program doesn't call itself again
		return
	} else {
		countdown(i - 1) // Recursion
	}
}

func fact(num int) int {
	if num == 1 {
		return 1		// Case case
	} else {
		return num * fact(num - 1)	// Recursion
	}
}

func main() {
	// countdown(10)
	fmt.Println(fact(5))
}