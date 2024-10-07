package log

import (
	"log"
	"os"
)

// InitLogger инициализирует стандартный логгер для игрового приложения.
func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
