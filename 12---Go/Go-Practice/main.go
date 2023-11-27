package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100000000; i++ {
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
