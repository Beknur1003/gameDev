package worker

import (
	"context"
	"testing"
	"time"
)

// TestProcessSlice проверяет корректность параллельной обработки слайса.
func TestProcessSlice(t *testing.T) {
	largeSlice := make([]int, 1_000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}

	// Установка контекста с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	totalSum := ProcessSlice(ctx, largeSlice, 10)

	// Ожидаемая сумма от 0 до 999 = (999 * 1000) / 2
	expectedSum := (999 * 1000) / 2

	if totalSum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, totalSum)
	}
}
