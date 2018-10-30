// Print duplicates lines found in Stdin
// or input file. If input files were provided
// then print file names where duplicates were found.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var dupFileNames map[string]bool
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupFileNames)
	} else {
		dupFileNames = make(map[string]bool)
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "exercise1.4: %v\n", err)
				continue
			}
			countLines(f, counts, dupFileNames)
			f.Close()
		}
	}
	if len(dupFileNames) != 0 {
		fmt.Printf("Duplicates were found in next files: ")
		var res, sep string
		for filename := range dupFileNames {
			res += sep + filename
			sep = " "
		}
		fmt.Println(res)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, duplFileNames map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if duplFileNames != nil && counts[line] > 1 {
			duplFileNames[f.Name()] = true
		}
	}
}
