package main

import (
	"fmt"
	"os"
)

func main() {
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	for index, arg := range os.Args {
		fmt.Printf("-%d\t%s\n", index, arg)
	}
}
