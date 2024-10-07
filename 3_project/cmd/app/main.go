package main

import (
	"context"
	"gameDev/3_project/internal/log"
	"gameDev/3_project/internal/worker"
	log2 "log"
	"time"
)

func main() {
	// Инициализация логгера
	log.InitLogger()

	// Создание большого слайса для обработки
	largeSlice := make([]int, 1_000_000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}

	// Число горутин для параллельной обработки
	numGoroutines := 10

	// Контекст с тайм-аутом для управления горутинами
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Параллельная обработка слайса
	totalSum := worker.ProcessSlice(ctx, largeSlice, numGoroutines)

	log2.Printf("Общая сумма элементов слайса: %d\n", totalSum)
	log2.Println("Все части обработаны.")
}
