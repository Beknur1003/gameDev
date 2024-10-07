package worker

import (
	"context"
	"log"
	"sync"
)

// ProcessSlice обрабатывает большой слайс, разбивая его на части и суммируя элементы параллельно.
func ProcessSlice(ctx context.Context, largeSlice []int, numGoroutines int) int {
	// Размер одной части для каждой горутины
	chunkSize := len(largeSlice) / numGoroutines
	var wg sync.WaitGroup

	// Канал для передачи результатов
	resultChannel := make(chan int, numGoroutines)
	wg.Add(numGoroutines)

	// Запуск горутин для обработки частей слайса
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(largeSlice) // Последняя горутина обрабатывает остаток слайса
		}

		go func(slicePart []int) {
			defer wg.Done()
			sum := 0
			for _, value := range slicePart {
				select {
				case <-ctx.Done():
					log.Println("Горутина прервана из-за отмены контекста")
					return
				default:
					sum += value
				}
			}
			resultChannel <- sum
		}(largeSlice[start:end])
	}

	// Закрытие канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Суммирование результатов
	totalSum := 0
	for partialSum := range resultChannel {
		totalSum += partialSum
	}

	return totalSum
}
