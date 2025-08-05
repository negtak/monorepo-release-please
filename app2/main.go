package main

import (
	"fmt"

	"github.com/negtak/monorepo-release-please/shared"
)

func main() {
	fmt.Println("App2 Starting...")
	
	msg := shared.FormatMessage("Hello from App2!")
	fmt.Println(msg)
	
	result := shared.Multiply(5, 8)
	fmt.Printf("5 × 8 = %d\n", result)
}