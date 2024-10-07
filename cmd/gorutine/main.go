package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(counter int) {
			time.Sleep(time.Second)
			fmt.Print(counter)

		}(i)
	}
	time.Sleep(2 * time.Second)
}
