package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() && len(input.Text()) != 0 {
		counts[input.Text()]++
	}

	for line, count := range counts {
		fmt.Printf("%s\t%d\n", line, count)
	}
}
