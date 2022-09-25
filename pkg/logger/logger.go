package logger

import (
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	log.Println("Initial Logger...")
	Debug = log.New(os.Stdout, "\033[32;1m[DEBUG]\033[0m ", log.LstdFlags)
	Info = log.New(os.Stdout, "\033[34;1m[INFO]\033[0m ", log.LstdFlags)
	Error = log.New(os.Stdout, "\033[31;1m[ERROR]\033[31;22m ", log.LstdFlags|log.Llongfile)
}
func Reset() string {
	return "\033[0m"
}
func Green() string {
	return "\033[32;1m"
}
func Blue() string {
	return "\033[34;1m"
}
func Red() string {
	return "\033[31;1m"
}
