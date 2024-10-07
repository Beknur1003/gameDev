package main

import (
	"gameDev/2_project/internal/handler"
	"gameDev/2_project/internal/log"
	"gameDev/2_project/internal/resource"
	log2 "log"
	"sync"
)

func main() {
	// Инициализация логгера
	log.InitLogger()

	// Создание игровых ресурсов
	resource1 := resource.NewResource("Game Resource 1")
	resource2 := resource.NewResource("Game Resource 2")

	var wg sync.WaitGroup
	wg.Add(2)

	// Запуск горутин для обработки игровых ресурсов
	go handler.SafeProcess(&wg, resource1, resource2)
	go handler.SafeProcess(&wg, resource2, resource1)

	wg.Wait()
	log2.Println("Все операции завершены успешно")
}
