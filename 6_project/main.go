package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Как и в 5_project, только если использовать атомарные операции для суммирования

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var totalSum int64
	var wg sync.WaitGroup

	numGoroutines := 2
	chunkSize := len(numbers) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(numbers)
		}

		go func(nums []int) {
			defer wg.Done()
			localSum := int64(0)
			for _, num := range nums {
				localSum += int64(num)
			}

			atomic.AddInt64(&totalSum, localSum)
		}(numbers[start:end])
	}

	wg.Wait()
	fmt.Println("Итоговая сумма:", totalSum)
}
