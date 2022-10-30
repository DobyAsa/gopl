package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	abort := make(chan struct{})

	go func(abort chan struct{}) {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}(abort)

	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("launch cancel!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Launch success!")
}
