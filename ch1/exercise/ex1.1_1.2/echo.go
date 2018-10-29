// Echo0 contains excersises 1.1 and 1.2 implementation
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println("The program name is", os.Args[0])
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
