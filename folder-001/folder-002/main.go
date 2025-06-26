package main

import (
	"fmt"
	"github.com/PeterCullenBurbery/go-functions" // the public module
)

func main() {
	gofunctions.SayHello("Peter")

	result := gofunctions.MultiplyBy2718(3)
	fmt.Println("3 × 2718 =", result)
}
