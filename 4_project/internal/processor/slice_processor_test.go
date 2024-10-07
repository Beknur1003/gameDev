package processor

import (
	"testing"
)

// TestProcessSlice проверяет корректность параллельной обработки слайса.
func TestProcessSlice(t *testing.T) {
	largeSlice := make([]int, 1_000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}

	// Создание процессора
	sliceProcessor := NewSliceProcessor()

	// Обработка слайса
	totalSum := sliceProcessor.ProcessSlice(largeSlice, 10)

	// Ожидаемая сумма от 0 до 999 = (999 * 1000) / 2
	expectedSum := (999 * 1000) / 2

	if totalSum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, totalSum)
	}
}
