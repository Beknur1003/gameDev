package processor

import (
	"sync"
)

// SliceProcessor отвечает за параллельную обработку большого слайса.
type SliceProcessor struct {
	mu       sync.Mutex
	totalSum int
	wg       sync.WaitGroup
}

// NewSliceProcessor создает новый экземпляр обработчика слайсов.
func NewSliceProcessor() *SliceProcessor {
	return &SliceProcessor{}
}

// ProcessSlice обрабатывает большой слайс параллельно, разбивая его на части.
func (p *SliceProcessor) ProcessSlice(largeSlice []int, numGoroutines int) int {
	chunkSize := len(largeSlice) / numGoroutines

	p.wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(largeSlice) // Последняя горутина обрабатывает остаток слайса
		}

		go func(slicePart []int) {
			defer p.wg.Done()
			sum := 0
			for _, value := range slicePart {
				sum += value
			}

			// Блокировка для безопасного доступа к общему состоянию
			p.mu.Lock()
			p.totalSum += sum
			p.mu.Unlock()
		}(largeSlice[start:end])
	}

	// Ожидание завершения всех горутин
	p.wg.Wait()

	return p.totalSum
}
