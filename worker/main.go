package main

import (
	"fmt"
	"time"
)

func main() {
	for _, i := range []int{1, 2, 3, 4, 5} {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Hello, World!")
}
