package main

import (
	"fmt"

	"github.com/negtak/monorepo-release-please/shared"
)

func main() {
	fmt.Println("App2 Starting...")

	msg := shared.FormatMessage("Hello from App2!!")
	fmt.Println(msg)

	result := shared.Multiply(9, 9)
	fmt.Printf("9 × 9 = %d\n", result)
}
