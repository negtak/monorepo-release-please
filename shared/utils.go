package shared

import (
	"fmt"
	"time"
)

func FormatMessage(msg string) string {
	return fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}