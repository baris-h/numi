package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(count int, min int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, count)
	
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}
	
	return numbers
}

func main() {
	count := 10
	min := 1
	max := 100
	
	randomNumbers := generateRandomNumbers(count, min, max)
	
	fmt.Println("Generated random numbers:", randomNumbers)
}
