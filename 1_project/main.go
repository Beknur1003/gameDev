package main

import (
	"fmt"
	"sync"
	"time"
)

// Пример кода: безопасный захват нескольких мьютексов

type Resource struct {
	name string
}

var resource1 = Resource{name: "Resource 1_project"}
var resource2 = Resource{name: "Resource 2"}

var mu1 = sync.Mutex{}
var mu2 = sync.Mutex{}

func process1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина process1: Попытка захватить mu1")
	mu1.Lock()
	fmt.Println("Горутина process1: mu1 захвачен")

	defer mu1.Unlock()

	time.Sleep(1 * time.Second)

	fmt.Println("Горутина process1: Попытка захватить mu2")
	mu2.Lock()
	fmt.Println("Горутина process1: mu2 захвачен")

	defer mu2.Unlock()

	fmt.Printf("Горутина process1: Обработка ресурсов %s и %s завершена\n", resource1.name, resource2.name)
}

func process2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина process2: Попытка захватить mu1")
	mu1.Lock()
	fmt.Println("Горутина process2: mu1 захвачен")

	defer mu1.Unlock()

	time.Sleep(1 * time.Second)

	fmt.Println("Горутина process2: Попытка захватить mu2")
	mu2.Lock()
	fmt.Println("Горутина process2: mu2 захвачен")

	defer mu2.Unlock()

	fmt.Printf("Горутина process2: Обработка ресурсов %s и %s завершена\n", resource1.name, resource2.name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go process1(&wg)
	go process2(&wg)

	wg.Wait()
	fmt.Println("Все операции завершены успешно")
}
