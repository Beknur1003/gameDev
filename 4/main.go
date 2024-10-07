package main

import (
	"fmt"
	"sync"
)

// как и в 3, но имеется обработка состояния гонки (race condition)

func main() {
	largeSlice := make([]int, 1_000_000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}

	numGoroutines := 10
	chunkSize := len(largeSlice) / numGoroutines

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalSum := 0

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(largeSlice)
		}

		go func(slicePart []int) {
			defer wg.Done()
			sum := 0
			for _, value := range slicePart {
				sum += value
			}

			mu.Lock()
			totalSum += sum
			mu.Unlock()
		}(largeSlice[start:end])
	}

	wg.Wait()
	fmt.Printf("Общая сумма элементов слайса: %d\n", totalSum)
	fmt.Println("Все части обработаны.")
}
