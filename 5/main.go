package main

import (
	"fmt"
	"sync"
)

// Можно заюзать мьютексы

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var totalSum int
	var wg sync.WaitGroup
	var mu sync.Mutex

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
			localSum := 0
			for _, num := range nums {
				localSum += num
			}

			mu.Lock()
			totalSum += localSum
			mu.Unlock()
		}(numbers[start:end])
	}

	wg.Wait()
	fmt.Println("Итоговая сумма:", totalSum)
}
