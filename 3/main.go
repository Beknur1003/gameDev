package main

import (
	"fmt"
	"sync"
)

// Оптимизированный пример распараллеливания работы с большим слайсом

func main() {
	largeSlice := make([]int, 1_000_000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}

	numGoroutines := 10
	chunkSize := len(largeSlice) / numGoroutines

	var wg sync.WaitGroup
	resultChannel := make(chan int, numGoroutines)

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
			resultChannel <- sum
		}(largeSlice[start:end])
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	totalSum := 0
	for partialSum := range resultChannel {
		totalSum += partialSum
	}

	fmt.Printf("Общая сумма элементов слайса: %d\n", totalSum)
	fmt.Println("Все части обработаны.")
}
