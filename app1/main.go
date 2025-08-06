package main

import (
	"fmt"

	"github.com/negtak/monorepo-release-please/shared"
)

func main() {
	fmt.Println("App1 Starting...")

	msg := shared.FormatMessage("Hello from App1!!")
	fmt.Println(msg)

	result := shared.Add(20, 20)
	fmt.Printf("20 + 20 = %d\n", result)
}
