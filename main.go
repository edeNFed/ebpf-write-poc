package main

import (
	"fmt"
	"time"
)

//go:noinline
func worker(str string) {
	headers := make(map[string]string)
	headers["X-Request-Id"] = str

	fmt.Printf("The Headers are: %s\n", headers)
}

func main() {
	for i := 0; i < 10; i++ {
		worker(fmt.Sprintf("request number: %d", i))
		time.Sleep(2 * time.Second)
	}
}
