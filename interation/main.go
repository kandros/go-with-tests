package main

import (
	"fmt"
)

func Repeat(characted string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += characted
	}
	return repeated
}

func main() {
	fmt.Println("Hello")
}
