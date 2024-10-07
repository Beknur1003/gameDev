package main

import (
	"gameDev/1_project/internal/handler"
	"gameDev/1_project/internal/log"
	"gameDev/1_project/internal/resource"
	"sync"
)

func main() {
	// Инициализация логгера
	log.InitLogger()

	// Создание ресурсов для игры
	resource1 := resource.NewResource("Game Resource 1")
	resource2 := resource.NewResource("Game Resource 2_project")

	var wg sync.WaitGroup
	wg.Add(2)

	// Запуск горутин для обработки игровых ресурсов
	go handler.ProcessResources(&wg, "GameProcess1", resource1, resource2)
	go handler.ProcessResources(&wg, "GameProcess2", resource2, resource1)

	wg.Wait()
}
