package resource

import (
	"log"
	"sync"
	"time"
)

// Resource представляет структуру игрового ресурса, используемую для блокировки.
type Resource struct {
	Name string
	mu   sync.Mutex
}

// NewResource создает новый игровой ресурс.
func NewResource(name string) *Resource {
	return &Resource{Name: name}
}

// Lock блокирует доступ к игровому ресурсу.
func (r *Resource) Lock() {
	r.mu.Lock()
	log.Printf("Игровой ресурс %s заблокирован", r.Name)
}

// Unlock разблокирует доступ к игровому ресурсу.
func (r *Resource) Unlock() {
	r.mu.Unlock()
	log.Printf("Игровой ресурс %s разблокирован", r.Name)
}

// Process симулирует обработку игрового ресурса.
func (r *Resource) Process() {
	time.Sleep(1 * time.Second) // Симуляция работы с игровым ресурсом
	log.Printf("Обработка игрового ресурса %s завершена", r.Name)
}
